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
	taskChain    chan int
	sharedClient *service.SvcImpl // 共享的客户端实例
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
		taskChain:    make(chan int, config.Thread),
		sharedClient: sharedClient,
	}
}

func (l *LoadGenerator) Start() {

	l.wg.Add(l.config.Thread)

	for i := 0; i < l.config.Thread; i++ {
		go l.worker(i)
	}
	go func() {
		for index := range l.taskChain {
			log.Infof("Restarting worker %d", index)
			l.wg.Add(1)
			go l.worker(index)
		}
	}()

	// Set up signal handling for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Wait for signal
	<-sigs
	log.Println("Received shutdown signal, stopping all goroutines...")

	// Cancel all goroutines
	l.cancel()

	// Wait for all goroutines to finish
	l.wg.Wait()

	// 调用清理函数
	l.sharedClient.CleanUp()

	log.Println("All goroutines stopped, exiting program.")
}

func (l *LoadGenerator) worker(index int) {
	defer func() {
		l.wg.Done()
	}()
	for {
		select {
		case <-l.ctx.Done():
			log.Infof("Goroutine %d exiting due to cancellation", index)
			return
		default:
			defer func() {
				if r := recover(); r != nil {
					buf := make([]byte, 1024)
					n := runtime.Stack(buf, false)
					stackTrace := string(buf[:n])
					log.Infof("Recovered from panic: %v\nStack trace:\n%s", r, stackTrace)
					l.taskChain <- index
				}
			}()

			chainCtx := NewContext(context.Background())
			chainCtx.Set(Client, l.sharedClient)
			start := time.Now()
			_, err := l.config.Chain.Execute(chainCtx)
			log.Infof("Thread %d executed chain, time used: %v", index, time.Since(start))
			if err != nil {
				log.Errorf("Error executing chain: %v", err)
			}
			if l.config.SleepTime > 0 {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(l.config.SleepTime)))
			}

		}
	}

}
