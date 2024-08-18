package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

func InputStartEndAndDate(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	QueryAll, err := cli.QueryStations()
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if QueryAll.Status != 1 {
		log.Errorf("Request failed, QueryAll.Status: %d, expected: %d", QueryAll.Status, 1)
		return nil, err
	}

	TheStartStationName, TheEndStationName, _ := randomlyChoosePlaces(QueryAll.Data)
	TheDepartureTime := extractDate(getRandomTime())
	//TheTrainTypeName :=

	ctx.Set(StartStation, TheStartStationName)
	ctx.Set(EndStation, TheEndStationName)
	ctx.Set(DepartureTime, TheDepartureTime)
	//ctx.Set(TrainTypeName, TheTrainTypeName)

	return nil, nil
}
