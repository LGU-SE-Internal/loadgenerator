package behaviors

import (
	"context"
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
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
	return data[key]
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
		result, err := node.Execute(ctx)
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
}

func (l *LoadGenerator) Start(conf ...func(*Config)) {
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

	var wg sync.WaitGroup
	wg.Add(config.Thread)

	for i := 0; i < config.Thread; i++ {
		go func(index int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					buf := make([]byte, 1024)
					n := runtime.Stack(buf, false)
					stackTrace := string(buf[:n])

					log.Printf("Recovered from panic: %v\nStack trace:\n%s", r, stackTrace)
				}
			}()

			for {
				ctx := NewContext(context.Background())
				ctx.Set(Client, service.NewSvcClients())
				_, err := config.Chain.Execute(ctx)
				if err != nil {
					log.Printf("Error executing chain: %v", err)
				}
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(config.SleepTime)))
			}
		}(i)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	wg.Wait()
}
