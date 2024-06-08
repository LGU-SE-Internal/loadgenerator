package behaviors

import (
	"math/rand"
	"sync"
)

var behaviors_ []BehaviorUnit
var once sync.Once

type Behavior interface {
	Run()
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
	config := Config{}
	for _, fn := range conf {
		fn(&config)
	}

	totalWeight := 0
	for _, behaviorUnit := range behaviors_ {
		totalWeight += behaviorUnit.Weight
	}

	randomWeight := rand.Intn(totalWeight)
	for _, behaviorUnit := range behaviors_ {
		randomWeight -= behaviorUnit.Weight
		if randomWeight <= 0 {
			behaviorUnit.B.Run()
			break
		}
	}
}
