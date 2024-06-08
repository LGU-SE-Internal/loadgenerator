package main

import "github.com/Lincyaw/loadgenerator/behaviors"

func main() {
	behaviors.RegisterBehaviors(
		behaviors.BehaviorUnit{
			B:      &behaviors.CreateUserBehavior{},
			Weight: 100,
		},
	)
	lg := &behaviors.LoadGenerator{}
	lg.Start(behaviors.WithThread(10))
}
