package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

func QueryPrice(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheRouteId := ctx.Get(RouteID).(string)
	TheTrainType := ctx.Get(TrainTypeName).(string)
	priceByRouteAndTrain, err := cli.FindByRouteIdAndTrainType(TheRouteId, TheTrainType)
	if err != nil {
		log.Errorf("FindByRouteIdAndTrainType failed: %v", err)
		return nil, err
	}
	if priceByRouteAndTrain.Status != 1 {
		log.Warnf("[Please change the traintype and try again] There is not corresponding Ticket available.")
		return &(NodeResult{false}), err // immediately end
	}

	ctx.Set(BasicPriceRate, priceByRouteAndTrain.Data.BasicPriceRate)
	ctx.Set(FirstClassPriceRate, priceByRouteAndTrain.Data.FirstClassPriceRate)

	return nil, nil
}
