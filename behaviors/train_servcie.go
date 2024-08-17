package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryTrain(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype all
	allTrainTypes, err := cli.QueryTraintype()
	if err != nil {
		log.Errorf("QueryTraintype all request failed, err %s", err)
		return nil, err
	}
	if allTrainTypes.Status != 1 {
		log.Errorf("allTrainTypes.Status != 1")
		return nil, err
	}
	if len(allTrainTypes.Data) == 0 {
		log.Errorf("QueryTraintype all returned no results")
		return nil, err
	}
	randomIndex := rand.Intn(len(allTrainTypes.Data))
	ctx.Set(TrainTypName, allTrainTypes.Data[randomIndex].Name)
	ctx.Set(ConfortClass, allTrainTypes.Data[randomIndex].ConfortClass)
	ctx.Set(AverageSpeed, allTrainTypes.Data[randomIndex].AverageSpeed)
	ctx.Set(EconomyClass, allTrainTypes.Data[randomIndex].EconomyClass)

	return nil, nil
}
