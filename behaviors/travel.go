package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
)

const ()

var TravelChain *Chain

func init() {
	TravelChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		return nil, nil
	}, "DummyTravelChain"))
	LoginChain.AddNextChain(NewChain(NewFuncNode(LoginAdmin, "LoginAdmin")), 1)
}

func QueryTripId(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.SeatService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	_ = cli

	ctx.Set(TripId, "D1345")
	return nil, nil
}
