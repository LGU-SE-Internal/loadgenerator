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
	"syscall"
	"time"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/Lincyaw/loadgenerator/stats"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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
	span := trace.SpanFromContext(ctx.ctx)

	result, err := f.fn(ctx)
	if err != nil && span.IsRecording() {
		span.RecordError(err)
		span.SetStatus(codes.Error, fmt.Sprintf("Error in node %s: %v", f.Name, err))
	}
	return result, err
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
			log.Tracef("Error occurred in node %s: %v", node.GetName(), err)
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
	config           *Config
	wg               sync.WaitGroup
	ctx              context.Context
	cancel           context.CancelFunc
	sharedClient     *service.SvcImpl // 共享的客户端实例
	printStatsTicker *time.Ticker
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

	sharedClient := service.NewSvcClients()

	return &LoadGenerator{
		config:           &config,
		ctx:              ctx,
		cancel:           cancel,
		sharedClient:     sharedClient,
		printStatsTicker: time.NewTicker(10 * time.Second),
	}
}

func (l *LoadGenerator) Start() {
	l.startStatsMonitor()

	l.wg.Add(l.config.Thread)
	for i := 0; i < l.config.Thread; i++ {
		go l.worker(i)
	}

	// Set up signal handling for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal
	<-sigs
	log.Println("Received shutdown signal, stopping all goroutines...")

	l.printStatsTicker.Stop()

	// Cancel the main context, which will cascade to all workers
	l.cancel()

	// Wait for all goroutines to finish with timeout
	done := make(chan struct{})
	go func() {
		l.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All workers stopped gracefully")
	case <-time.After(30 * time.Second):
		log.Warn("Timeout waiting for workers to stop, forcing exit")
	}

	l.sharedClient.CleanUp()

	runtime.GC()

	log.Println("All goroutines stopped, exiting program.")
}

func (l *LoadGenerator) worker(index int) {
	defer l.wg.Done()

	for {
		select {
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
					}
				}()

				chainCtx := NewContext(context.Background())
				chainCtx.Set(Client, l.sharedClient)
				start := time.Now()
				_, err := l.config.Chain.Execute(chainCtx)
				if time.Since(start) > 5*time.Second {
					log.Errorf("Thread %d executed chain, time used: %v", index, time.Since(start))
				}
				if err != nil {
					log.Warn(err)
				}

				if l.config.SleepTime > 0 {
					time.Sleep(time.Millisecond * time.Duration(l.config.SleepTime))
				}
			}()
		}
	}
}

func (l *LoadGenerator) startStatsMonitor() {
	log.Info("Starting stats monitor...")
	cleanupTicker := time.NewTicker(5 * time.Minute)
	gcTicker := time.NewTicker(2 * time.Minute)

	go func() {
		defer cleanupTicker.Stop()
		defer gcTicker.Stop()
		for {
			select {
			case <-l.ctx.Done():
				return
			case <-l.printStatsTicker.C:
				l.printLatencyStats()
			case <-cleanupTicker.C:
				log.Info("Cleaning old statistics records...")
				stats.GlobalLatencyManager.CleanOldRecords(10 * time.Minute)
			case <-gcTicker.C:
				log.Debug("Forcing garbage collection...")
				runtime.GC()
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				log.Errorf("Memory usage: Alloc=%d KB, TotalAlloc=%d KB, Sys=%d KB, NumGC=%d",
					m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
			}
		}
	}()
}

func (l *LoadGenerator) printLatencyStats() {
	topSlowStats := stats.GlobalLatencyManager.GetTopSlowStats(10)

	log.Warn("=== Top 10 Slowest Requests ===")
	if len(topSlowStats) == 0 {
		log.Info("No request statistics available yet")
	} else {
		for i, statObj := range topSlowStats {
			min, max, avg, p50, p95, p99 := statObj.GetStats()
			log.Warnf("#%d URL: %s %s", i+1, statObj.Method, statObj.URL)
			log.Warnf("  Count: %d", statObj.Count)
			log.Warnf("  Min: %v, Max: %v, Avg: %v", min, max, avg)
			log.Warnf("  P50: %v, P95: %v, P99: %v", p50, p95, p99)
			log.Warnf("---")
		}
	}

	log.Warnf("Current Threads: %d, Sleep Time: %d ms", l.config.Thread, l.config.SleepTime)
	log.Warn("=========================")
}
