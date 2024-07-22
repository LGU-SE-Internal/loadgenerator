package main

import (
	"github.com/Lincyaw/loadgenerator/behaviors"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	behaviors.RegisterBehaviors(
		behaviors.BehaviorUnit{
			B:      &behaviors.CreateUserBehavior{},
			Weight: 100,
		},
		behaviors.BehaviorUnit{
			B:      &behaviors.PreserveBehavior{},
			Weight: 100,
		},
		behaviors.BehaviorUnit{
			B:      &behaviors.TravelplanBehavior{},
			Weight: 100,
		},
		behaviors.BehaviorUnit{
			B:      &behaviors.TravelBehavior{},
			Weight: 100,
		},
	)
	lg := &behaviors.LoadGenerator{}
	lg.Start(behaviors.WithThread(1), behaviors.WithSleep(1000))
}
