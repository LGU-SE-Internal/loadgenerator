package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AllRoutesByQuery, err := cli.QueryAllRoutes()
	if err != nil {
		log.Errorf("Request failed, err2 %s", err)
		return nil, err
	}
	if AllRoutesByQuery.Status != 1 {
		log.Fatal("AllRoutes_By_Query.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(AllRoutesByQuery.Data))
	debug := AllRoutesByQuery.Data[randomIndex]
	fmt.Println(debug)
	ctx.Set(RouteID, AllRoutesByQuery.Data[randomIndex].Id)
	ctx.Set(StartStation, AllRoutesByQuery.Data[randomIndex].StartStation)
	ctx.Set(EndStation, AllRoutesByQuery.Data[randomIndex].EndStation)
	ctx.Set(StationName, AllRoutesByQuery.Data[randomIndex].Stations)
	ctx.Set(Distances, AllRoutesByQuery.Data[randomIndex].Distances)

	return nil, nil
}
