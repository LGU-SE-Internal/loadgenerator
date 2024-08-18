package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryTripInfo(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheTrainTypeName := ctx.Get(TrainTypeName).(string)
	log.Infof("The train type is %s", TheTrainTypeName)

	tripInfo := service.TripInfo{
		StartPlace:    ctx.Get(StartStation).(string),
		EndPlace:      ctx.Get(EndStation).(string),
		DepartureTime: ctx.Get(DepartureTime).(string),
	}

	//var queryInfoResp service.QueryInfoResponse

	switch TheTrainTypeName {
	case "GaoTieOne", "GaoTieTwo", "DongCheOne": // travel

		var travelSvc service.TravelService = cli
		queryInfoResp, err := travelSvc.QueryInfo(tripInfo)
		if err != nil {
			log.Errorf("QueryInfo request failed, err %s", err)
			return nil, err
		}
		if queryInfoResp.Status != 1 {
			log.Errorf("QueryInfo failed, status: %d", queryInfoResp.Status)
			return nil, err
		}

		if len(queryInfoResp.Data) == 0 {
			log.Println("[Please Select Other Train type For The Current Start_End Pair] QueryInfo response is empty.")
			return &(NodeResult{false}), nil
		}

		randomIndex := rand.Intn(len(queryInfoResp.Data))
		ctx.Set(TripID, fmt.Sprintf("%s%s", queryInfoResp.Data[randomIndex].TripId.Type, queryInfoResp.Data[randomIndex].TripId.Number))
		//ctx.Set(TrainTypeName, queryInfoResp.Data[randomIndex].TrainTypeName)
		/*	ctx.Set(StartStation, queryInfoResp.Data[randomIndex].StartStation)
			ctx.Set(TerminalStation, queryInfoResp.Data[randomIndex].TerminalStation)*/
		ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
		ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
		ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
		ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
		ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
		ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)

	default: // travel2

		var travel2Svc service.Travel2Service = cli
		queryInfoResp, err := travel2Svc.QueryByBatch(&tripInfo)
		if err != nil {
			log.Errorf("QueryInfo request failed, err %s", err)
			return nil, err
		}
		if queryInfoResp.Status != 1 {
			log.Errorf("QueryInfo failed, status: %d", queryInfoResp.Status)
			return nil, err
		}

		if len(queryInfoResp.Data) == 0 {
			log.Println("[Please Select Other Train type For The Current Start_End Pair] QueryInfo response is empty.")
			return &(NodeResult{false}), nil
		}

		randomIndex := rand.Intn(len(queryInfoResp.Data))
		ctx.Set(TripID, fmt.Sprintf("%s%s", queryInfoResp.Data[randomIndex].TripId.Type, queryInfoResp.Data[randomIndex].TripId.Number))
		//ctx.Set(TrainTypeName, queryInfoResp.Data[randomIndex].TrainTypeName)
		/*	ctx.Set(StartStation, queryInfoResp.Data[randomIndex].StartStation)
			ctx.Set(TerminalStation, queryInfoResp.Data[randomIndex].TerminalStation)*/
		ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
		ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
		ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
		ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
		ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
		ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)
	}

	//randomIndex := rand.Intn(len(queryInfoResp.Data))
	//ctx.Set(TripID, fmt.Sprintf("%s%s", queryInfoResp.Data[randomIndex].TripId.Type, queryInfoResp.Data[randomIndex].TripId.Number))
	////ctx.Set(TrainTypeName, queryInfoResp.Data[randomIndex].TrainTypeName)
	///*	ctx.Set(StartStation, queryInfoResp.Data[randomIndex].StartStation)
	//	ctx.Set(TerminalStation, queryInfoResp.Data[randomIndex].TerminalStation)*/
	//ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	//ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	//ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	//ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	//ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	//ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)

	return nil, nil
}
