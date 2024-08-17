package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
)

func QueryTripInfo(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.Travel2Service)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	info := service.TripInfo{
		StartPlace:    ctx.Get(StartStation).(string),
		EndPlace:      ctx.Get(EndStation).(string),
		DepartureTime: getRandomTime(),
	}
	batch, err := cli.QueryByBatch(&info)
	if err != nil {
		return nil, err
	}
	_ = batch
	return nil, nil
}
