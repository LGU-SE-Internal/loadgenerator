package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryStationFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllStationFood()
	if err != nil {
		log.Errorf("Resp returns err: %v", err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Errorf("GetAllStationFood status should be 1, but is %d", resp.Status)
		return nil, err
	}
	randomIndex := rand.Intn(len(resp.Data))
	ctx.Set(StoreName, resp.Data[randomIndex].StoreName)
	ctx.Set(Phone, resp.Data[randomIndex].Telephone)
	ctx.Set(Price, resp.Data[randomIndex].DeliveryFee)

	return nil, nil

}
