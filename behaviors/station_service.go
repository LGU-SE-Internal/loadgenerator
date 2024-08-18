package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	QueryAll, err7 := cli.QueryStations()
	if err7 != nil {
		log.Errorf("Request failed, err7 %s", err7)
		return nil, err7
	}
	if QueryAll.Status != 1 {
		log.Errorf("Request failed, QueryAll.Status: %d, expected: %d", QueryAll.Status, 1)
		return nil, err7
	}
	randomIndex := rand.Intn(len(QueryAll.Data))
	ctx.Set(StationId, QueryAll.Data[randomIndex].Id)
	ctx.Set(StationNames, QueryAll.Data[randomIndex].Name)
	ctx.Set(StayTime, QueryAll.Data[randomIndex].StayTime)

	return nil, nil
}
