package behaviors

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/Lincyaw/loadgenerator/stats"
	log "github.com/sirupsen/logrus"
)

const Client = "client"

type ContextKey string

const dataKey = ContextKey("data")

// Context wraps context.Context and provides additional methods
type Context struct {
	ctx context.Context
}

func NewContext(ctx context.Context) *Context {
	return &Context{ctx: ctx}
}

// Set sets a value in the context
func (c *Context) Set(key string, value interface{}) {
	data := c.getDataMap()
	data[key] = value
	c.ctx = context.WithValue(c.ctx, dataKey, data)
}

// Get retrieves a value from the context
func (c *Context) Get(key string) interface{} {
	data := c.getDataMap()
	value, ok := data[key]
	if !ok {
		log.Errorf("There is no value for key %s", key)
	}
	return value
}

func (c *Context) getDataMap() map[string]interface{} {
	data, ok := c.ctx.Value(dataKey).(map[string]interface{})
	if !ok {
		data = make(map[string]interface{})
	}
	return data
}

// Node represents a single node in the chain
type Node interface {
	Execute(ctx *Context) (*NodeResult, error)
	GetName() string
}

type NodeResult struct {
	Continue bool
}

type FuncNode struct {
	fn   func(*Context) (*NodeResult, error)
	Name string
}

func (f *FuncNode) Execute(ctx *Context) (*NodeResult, error) {
	return f.fn(ctx)
}

func (f *FuncNode) GetName() string {
	return f.Name
}

func NewFuncNode(fn func(*Context) (*NodeResult, error), name string) *FuncNode {
	return &FuncNode{fn: fn, Name: name}
}

type Chain struct {
	nodes          []Node
	nextChains     []chainWithProbability
	probabilitySum float64
	Name           string
}
type chainWithProbability struct {
	chain       *Chain
	probability float64
}

func NewChain(nodes ...Node) *Chain {
	return &Chain{nodes: nodes}
}

func (c *Chain) AddNode(node Node) {
	c.nodes = append(c.nodes, node)
}

func (c *Chain) AddNextChain(next *Chain, probability float64) {
	c.nextChains = append(c.nextChains, chainWithProbability{chain: next, probability: probability})
	c.probabilitySum += probability
}

func (c *Chain) Execute(ctx *Context) (*NodeResult, error) {
	for _, node := range c.nodes {
		startT := time.Now()
		result, err := node.Execute(ctx)
		log.Debugf("Executed node %s, time used: %v", node.GetName(), time.Since(startT))
		if err != nil {
			return nil, err
		}
		if result == nil {
			continue
		}
		if !result.Continue {
			return nil, nil
		}
	}

	if len(c.nextChains) > 0 {
		randValue := rand.Float64() * c.probabilitySum
		cumulative := 0.0
		for _, cp := range c.nextChains {
			cumulative += cp.probability
			if randValue <= cumulative {
				return cp.chain.Execute(ctx)
			}
		}
	}

	return nil, nil
}

func (c *Chain) GetName() string {
	return c.Name
}

func (c *Chain) VisualizeChain(level int) string {
	result := ""

	// 打印当前链的节点
	for _, node := range c.nodes {
		result += fmt.Sprintf("%sNode: %s\n", getIndent(level), node.GetName())
	}

	// 打印下一级链的信息
	for _, nextChain := range c.nextChains {
		result += fmt.Sprintf("%sProbability: %.2f\n", getIndent(level), nextChain.probability)
		result += nextChain.chain.VisualizeChain(level + 1)
	}

	return result
}

func getIndent(level int) string {
	return "  " + strings.Repeat("  ", level)
}

type Config struct {
	Thread    int
	SleepTime int
	Chain     *Chain
}

func WithThread(thread int) func(*Config) {
	return func(conf *Config) {
		conf.Thread = thread
	}
}
func WithSleep(milli int) func(*Config) {
	return func(conf *Config) {
		conf.SleepTime = milli
	}
}
func WithChain(c *Chain) func(*Config) {
	return func(conf *Config) {
		conf.Chain = c
	}
}

type LoadGenerator struct {
	config       *Config
	wg           sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
	sharedClient *service.SvcImpl // 共享的客户端实例

	// 动态负载控制
	currentThreads       int32
	currentSleepTime     int32
	maxThreads           int32
	minThreads           int32
	statsCheckTicker     *time.Ticker
	printStatsTicker     *time.Ticker
	slowRequestThreshold time.Duration

	// 用于管理 worker 线程
	workerContexts map[int]context.Context
	workerCancels  map[int]context.CancelFunc
	workerMutex    sync.RWMutex
	nextWorkerID   int32
}

func NewLoadGenerator(conf ...func(*Config)) *LoadGenerator {
	ctx, cancel := context.WithCancel(context.Background())
	config := Config{}
	for _, fn := range conf {
		fn(&config)
	}

	if config.Thread <= 0 {
		config.Thread = 1
	}

	if config.Chain == nil {
		panic("LoadGenerator needs chain")
	}

	// 初始化共享的客户端实例
	sharedClient := service.NewSvcClients()

	return &LoadGenerator{
		config:       &config,
		ctx:          ctx,
		cancel:       cancel,
		sharedClient: sharedClient,

		// 动态负载控制初始化
		currentThreads:       int32(config.Thread),
		currentSleepTime:     int32(config.SleepTime),
		maxThreads:           int32(config.Thread * 2), // 最大线程数为初始值的2倍
		minThreads:           1,
		statsCheckTicker:     time.NewTicker(30 * time.Second), // 每30秒检查一次
		printStatsTicker:     time.NewTicker(10 * time.Second), // 每10秒打印一次统计信息
		slowRequestThreshold: 10 * time.Second,

		// 初始化 worker 管理
		workerContexts: make(map[int]context.Context),
		workerCancels:  make(map[int]context.CancelFunc),
		nextWorkerID:   0,
	}
}

func (l *LoadGenerator) Start() {
	l.startStatsMonitor()

	currentThreads := atomic.LoadInt32(&l.currentThreads)
	l.wg.Add(int(currentThreads))

	for i := 0; i < int(currentThreads); i++ {
		l.startWorker(-1) // 使用-1表示自动分配ID
	}

	// Set up signal handling for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal
	<-sigs
	log.Println("Received shutdown signal, stopping all goroutines...")

	l.statsCheckTicker.Stop()
	l.printStatsTicker.Stop()

	// Cancel all goroutines
	l.cancel()

	// Cancel all worker contexts
	l.workerMutex.Lock()
	for _, cancel := range l.workerCancels {
		cancel()
	}
	l.workerMutex.Unlock()

	// Wait for all goroutines to finish
	l.wg.Wait()

	// 调用清理函数
	l.sharedClient.CleanUp()

	log.Println("All goroutines stopped, exiting program.")
}

func (l *LoadGenerator) worker(index int, workerCtx context.Context) {
	defer func() {
		l.wg.Done()
		// 清理 worker context
		l.workerMutex.Lock()
		delete(l.workerContexts, index)
		delete(l.workerCancels, index)
		l.workerMutex.Unlock()
		log.Infof("Worker %d cleanup completed", index)
	}()

	for {
		select {
		case <-workerCtx.Done():
			log.Infof("Worker %d exiting due to cancellation", index)
			return
		case <-l.ctx.Done():
			log.Infof("Worker %d exiting due to main context cancellation", index)
			return
		default:
			func() {
				defer func() {
					if r := recover(); r != nil {
						buf := make([]byte, 1024)
						n := runtime.Stack(buf, false)
						stackTrace := string(buf[:n])
						log.Errorf("Worker %d recovered from panic: %v\nStack trace:\n%s", index, r, stackTrace)
						// 不要在这里重启worker，让它自然退出
						// 动态负载控制会在必要时创建新的worker
					}
				}()

				chainCtx := NewContext(context.Background())
				chainCtx.Set(Client, l.sharedClient)
				start := time.Now()
				_, err := l.config.Chain.Execute(chainCtx)
				log.Infof("Thread %d executed chain, time used: %v", index, time.Since(start))
				if err != nil {
					log.Warn(err)
				}

				// 使用动态睡眠时间
				currentSleepTime := atomic.LoadInt32(&l.currentSleepTime)
				if currentSleepTime > 0 {
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(int(currentSleepTime))))
				}
			}()
		}
	}
}

// startWorker 启动一个新的 worker 并管理其 context
// 如果index为-1，则自动分配一个新的ID
func (l *LoadGenerator) startWorker(index int) {
	var workerID int

	if index == -1 {
		// 自动分配新的ID
		workerID = int(atomic.AddInt32(&l.nextWorkerID, 1))
	} else {
		// 使用指定的ID（通常用于重启）
		workerID = index
	}

	workerCtx, workerCancel := context.WithCancel(l.ctx)

	l.workerMutex.Lock()
	l.workerContexts[workerID] = workerCtx
	l.workerCancels[workerID] = workerCancel
	l.workerMutex.Unlock()

	go l.worker(workerID, workerCtx)
}

// stopWorker 停止指定的 worker
func (l *LoadGenerator) stopWorker(index int) {
	l.workerMutex.Lock()
	defer l.workerMutex.Unlock()

	if cancel, exists := l.workerCancels[index]; exists {
		cancel()
		delete(l.workerContexts, index)
		delete(l.workerCancels, index)
		log.Infof("Stopped worker %d", index)
	}
}

// getActiveWorkerCount 获取当前活跃的worker数量
func (l *LoadGenerator) getActiveWorkerCount() int {
	l.workerMutex.RLock()
	defer l.workerMutex.RUnlock()
	return len(l.workerContexts)
}

// 动态负载控制方法
func (l *LoadGenerator) adjustLoadBasedOnStats() {
	// 需要导入httpclient包来访问GlobalLatencyManager
	// 这里我们先检查是否有超过阈值的请求
	hasSlowRequests := l.checkSlowRequests()

	if hasSlowRequests {
		l.decreaseLoad()
	} else {
		l.increaseLoad()
	}
}

func (l *LoadGenerator) checkSlowRequests() bool {
	return stats.GlobalLatencyManager.HasSlowRequests(l.slowRequestThreshold)
}

func (l *LoadGenerator) decreaseLoad() {
	currentThreads := atomic.LoadInt32(&l.currentThreads)
	activeWorkers := l.getActiveWorkerCount()
	currentSleepTime := atomic.LoadInt32(&l.currentSleepTime)

	// 减少线程数或增加睡眠时间
	if currentThreads > l.minThreads && activeWorkers > 0 {
		newThreads := currentThreads - 1
		atomic.StoreInt32(&l.currentThreads, newThreads)

		// 停止一个 worker
		l.workerMutex.RLock()
		var workerToStop int = -1
		for workerIndex := range l.workerCancels {
			workerToStop = workerIndex
			break
		}
		l.workerMutex.RUnlock()

		if workerToStop >= 0 {
			l.stopWorker(workerToStop)
		}

		log.Infof("Decreased threads from %d to %d due to slow requests (active workers: %d)", currentThreads, newThreads, activeWorkers-1)
	} else {
		// 如果线程数已经是最小值，则增加睡眠时间
		newSleepTime := currentSleepTime + 1000 // 增加1秒
		atomic.StoreInt32(&l.currentSleepTime, newSleepTime)
		log.Infof("Increased sleep time from %d to %d ms due to slow requests", currentSleepTime, newSleepTime)
	}
}

func (l *LoadGenerator) increaseLoad() {
	currentThreads := atomic.LoadInt32(&l.currentThreads)
	currentSleepTime := atomic.LoadInt32(&l.currentSleepTime)

	// 如果睡眠时间大于原始配置，先减少睡眠时间
	if currentSleepTime > int32(l.config.SleepTime) {
		newSleepTime := currentSleepTime - 500 // 减少500ms
		if newSleepTime < int32(l.config.SleepTime) {
			newSleepTime = int32(l.config.SleepTime)
		}
		atomic.StoreInt32(&l.currentSleepTime, newSleepTime)
		log.Infof("Decreased sleep time from %d to %d ms", currentSleepTime, newSleepTime)
	} else if currentThreads < l.maxThreads {
		// 如果睡眠时间已经是原始值，则增加线程数
		newThreads := currentThreads + 1
		atomic.StoreInt32(&l.currentThreads, newThreads)
		log.Infof("Increased threads from %d to %d", currentThreads, newThreads)

		// 启动新的worker
		l.wg.Add(1)
		l.startWorker(-1) // 使用-1表示自动分配ID
	}
}

func (l *LoadGenerator) startStatsMonitor() {
	log.Info("Starting stats monitor...")
	go func() {
		for {
			select {
			case <-l.ctx.Done():
				return
			case <-l.statsCheckTicker.C:
				l.adjustLoadBasedOnStats()
			case <-l.printStatsTicker.C:
				l.printLatencyStats()
			}
		}
	}()
}

func (l *LoadGenerator) printLatencyStats() {
	// 获取延迟最高的前10个请求
	topSlowStats := stats.GlobalLatencyManager.GetTopSlowStats(10)

	log.Info("=== Top 10 Slowest Requests ===")
	if len(topSlowStats) == 0 {
		log.Info("No request statistics available yet")
	} else {
		for i, statObj := range topSlowStats {
			min, max, avg, p50, p95, p99 := statObj.GetStats()
			log.Infof("#%d URL: %s %s", i+1, statObj.Method, statObj.URL)
			log.Infof("  Count: %d", statObj.Count)
			log.Infof("  Min: %v, Max: %v, Avg: %v", min, max, avg)
			log.Infof("  P50: %v, P95: %v, P99: %v", p50, p95, p99)
			log.Info("---")
		}
	}

	currentThreads := atomic.LoadInt32(&l.currentThreads)
	activeWorkers := l.getActiveWorkerCount()
	currentSleepTime := atomic.LoadInt32(&l.currentSleepTime)
	slowRequestsCount := stats.GlobalLatencyManager.GetSlowRequestsCount(l.slowRequestThreshold)

	log.Infof("Current Threads: %d, Active Workers: %d, Sleep Time: %d ms", currentThreads, activeWorkers, currentSleepTime)
	log.Infof("Slow requests (>10s): %d", slowRequestsCount)
	log.Infof("Next Worker ID: %d", atomic.LoadInt32(&l.nextWorkerID))
	log.Info("=========================")
}
