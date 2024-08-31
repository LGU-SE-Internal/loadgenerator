package main

import (
	"github.com/Lincyaw/loadgenerator/behaviors"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lg := &behaviors.LoadGenerator{}
	lg.Start(behaviors.WithThread(1), behaviors.WithSleep(1000), behaviors.WithChain(behaviors.OrderConsignChain))
}
