package main

import "github.com/Lincyaw/loadgenerator/behaviors"

func main() {
	behaviors.RegisterBehaviors(
		behaviors.BehaviorUnit{
			B:      &behaviors.OrderTicketsBehavior{},
			Weight: 100,
		},
	)

}
