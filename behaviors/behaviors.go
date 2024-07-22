package behaviors

import (
	"github.com/Lincyaw/loadgenerator/service"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var behaviors_ []BehaviorUnit
var once sync.Once

type Behavior interface {
	Run(cli *service.SvcImpl)
}

type BehaviorUnit struct {
	B      Behavior
	Weight int
}

func RegisterBehaviors(behaviors ...BehaviorUnit) {
	once.Do(func() {
		behaviors_ = make([]BehaviorUnit, 0)
	})
	for _, behavior := range behaviors {
		behaviors_ = append(behaviors_, behavior)
	}
}

func GetBehaviors() []BehaviorUnit {
	return behaviors_
}

type Config struct {
	Thread    int
	SleepTime int
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

	totalWeight := 0
	weightBoundaries := make([]int, len(behaviors_))

	for i, behaviorUnit := range behaviors_ {
		totalWeight += behaviorUnit.Weight
		weightBoundaries[i] = totalWeight
	}

	var wg sync.WaitGroup
	wg.Add(config.Thread)

	for i := 0; i < config.Thread; i++ {
		go func(index int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					// 处理异常，比如记录日志
					buf := make([]byte, 1024)
					n := runtime.Stack(buf, false)
					stackTrace := string(buf[:n])

					// 记录日志，包括 panic 信息和调用栈
					log.Printf("Recovered from panic: %v\nStack trace:\n%s", r, stackTrace)
				}
			}()

			randSrc := rand.NewSource(time.Now().UnixNano())
			randGen := rand.New(randSrc)

			for {
				randomWeight := randGen.Intn(totalWeight)
				selectedIndex := 0
				for j, boundary := range weightBoundaries {
					if randomWeight < boundary {
						selectedIndex = j
						break
					}
				}
				behaviors_[selectedIndex].B.Run(service.NewSvcClients())
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
