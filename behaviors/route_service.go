package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

func QueryRouteByStartAndEnd(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheStart := ctx.Get(StartStation).(string)
	TheEnd := ctx.Get(EndStation).(string)
	var routeSvc service.RouteService = cli
	AllRoutesByQueryStartAndEnd, err := routeSvc.QueryRoutesByStartAndEnd(TheStart, TheEnd)
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if AllRoutesByQueryStartAndEnd.Status != 1 { // With Prob = (156-10) / 156 approximately equivalent to 94%
		log.Infof("Can not find the Start-End pair, query parameter, start:[%s], end[%s]", TheStart, TheEnd)
		return &(NodeResult{false}), err // immediately end
	}

	if len(AllRoutesByQueryStartAndEnd.Data) == 0 {
		log.Warnf("No routes found for Start-End pair, start:[%s], end[%s]", TheStart, TheEnd)
		return &(NodeResult{false}), nil
	}

	randomIndex := rand.Intn(len(AllRoutesByQueryStartAndEnd.Data))
	ctx.Set(RouteID, AllRoutesByQueryStartAndEnd.Data[randomIndex].Id)
	/*	ctx.Set(StartStation, AllRoutesByQueryStartAndEnd.Data[randomIndex].StartStation)
		ctx.Set(EndStation, AllRoutesByQueryStartAndEnd.Data[randomIndex].EndStation)*/
	ctx.Set(StationName, AllRoutesByQueryStartAndEnd.Data[randomIndex].Stations)
	ctx.Set(Distances, AllRoutesByQueryStartAndEnd.Data[randomIndex].Distances)

	return nil, nil
}

func ChooseRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	var routeSvc service.RouteService = cli
	allRoutes, err := routeSvc.QueryAllRoutes()
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if allRoutes.Status != 1 {
		log.Errorf("queryAllRoutes failed, data: %+v", allRoutes)
		return &(NodeResult{false}), err
	}

	if len(allRoutes.Data) == 0 {
		log.Warnf("No routes found")
		return &(NodeResult{false}), nil
	}

	randomIndex := rand.Intn(len(allRoutes.Data))
	TheDepartureTime := extractDate(getRandomTime())
	ctx.Set(RouteID, allRoutes.Data[randomIndex].Id)
	ctx.Set(StartStation, allRoutes.Data[randomIndex].StartStation)
	ctx.Set(EndStation, allRoutes.Data[randomIndex].EndStation)
	ctx.Set(DepartureTime, TheDepartureTime)
	ctx.Set(StationName, allRoutes.Data[randomIndex].Stations)
	ctx.Set(Distances, allRoutes.Data[randomIndex].Distances)
	return nil, nil
}
