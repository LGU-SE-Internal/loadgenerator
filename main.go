package main

import (
	"github.com/Lincyaw/loadgenerator/behaviors"
)

func main() {
	lg := &behaviors.LoadGenerator{}
	composedChain := behaviors.NewChain(behaviors.NewFuncNode(func(ctx *behaviors.Context) (*behaviors.NodeResult, error) {
		return nil, nil
	}, "dummy"))
	composedChain.AddNextChain(behaviors.OrderChangeChain, 0.5)
	//composedChain.AddNextChain(behaviors.NormalOrderPayChain, 0.5)
	lg.Start(behaviors.WithThread(1), behaviors.WithSleep(1000), behaviors.WithChain(composedChain))
}
