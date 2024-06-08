package behaviors

import (
	"github.com/Lincyaw/loadgenerator/service"
	"math/rand"
	"sync"
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
	Thread int
}

func WithThread(thread int) func(*Config) {
	return func(conf *Config) {
		conf.Thread = thread
	}
}

type LoadGenerator struct {
}

func (l *LoadGenerator) Start(conf ...func(*Config)) {
	cli := service.NewSvcClients()
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
		go func() {
			defer wg.Done()
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
				behaviors_[selectedIndex].B.Run(cli)
			}
		}()
	}

	cli.ShowStats()
}
