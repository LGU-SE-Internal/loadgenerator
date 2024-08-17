package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func CreateSeat(ctx *Context) (*NodeResult, error) {
	// cli, ok := ctx.Get(Client).(*service.SvcImpl)
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	totalNum := rand.Intn(10) + 10
	seatCreateInfoReq := &service.SeatCreateInfoReq{
		TravelDate:  ctx.Get(TravelDate).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		DestStation: ctx.Get(EndStation).(string),
		SeatType:    ctx.Get(SeatClass).(int),
		TotalNum:    totalNum,
		Stations:    ctx.Get(StationName).([]string),
	}

	resp, err := cli.ReqSeatCreate(seatCreateInfoReq)

	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Errorf("SeatCreateInfoReq.Status != 1, resp: %+v", resp)
		return nil, err
	}

	ctx.Set(SeatNo, resp.Data.SeatNo)
	ctx.Set(DestStation, resp.Data.DestStation)
	//ctx.Set(StartStation, resp.Data.StartStation)

	return nil, nil
}

func QuerySeatInfo(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.SeatService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	_ = cli

	ctx.Set(SeatClass, 2)
	return nil, nil
}
