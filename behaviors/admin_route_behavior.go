package behaviors

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// Admin Route Behavior Chain - 管理路线的增删查操作
var AdminRouteChain *Chain

func init() {
	AdminRouteChain = NewChain(
		NewFuncNode(LoginAdmin, "LoginAdmin"),
		NewFuncNode(AdminQueryAllRoutes, "AdminQueryAllRoutes"),
	)

	// 添加后续的路线管理操作链
	AdminRouteChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddRoute, "AdminAddRoute"),
	), 0.6)

	AdminRouteChain.AddNextChain(NewChain(
		NewFuncNode(AdminDeleteRoute, "AdminDeleteRoute"),
	), 0.4)
}

// AdminQueryAllRoutes 查询所有路线
func AdminQueryAllRoutes(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.ReqGetAllRoutes()
	if err != nil {
		log.Errorf("AdminQueryAllRoutes failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminQueryAllRoutes returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	if len(resp.Data) > 0 {
		randomIndex := rand.Intn(len(resp.Data))
		ctx.Set(RouteID, resp.Data[randomIndex].ID)
		ctx.Set(StartStation, resp.Data[randomIndex].StartStation)
		ctx.Set(EndStation, resp.Data[randomIndex].EndStation)
		ctx.Set(StationName, strings.Split(resp.Data[randomIndex].StationList, ","))
		ctx.Set(Distances, resp.Data[randomIndex].DistanceList)
	}

	log.Infof("AdminQueryAllRoutes returned %d routes", len(resp.Data))
	return nil, nil
}

// AdminAddRoute 添加路线
func AdminAddRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// 生成随机的车站列表
	stations := generateRandomStationList()
	distances := generateRandomDistances(len(stations))

	route := &service.AdminRouteInfo{
		ID:           uuid.New().String(),
		LoginID:      ctx.Get(LoginToken).(string),
		StartStation: stations[0],
		EndStation:   stations[len(stations)-1],
		StationList:  strings.Join(stations, ","),
		DistanceList: distances,
	}

	resp, err := cli.ReqAddRoute(route)
	if err != nil {
		log.Errorf("AdminAddRoute failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddRoute returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(RouteID, route.ID)
	ctx.Set(StartStation, route.StartStation)
	ctx.Set(EndStation, route.EndStation)
	ctx.Set(StationName, stations)

	log.Infof("AdminAddRoute success: routeId=%s, from %s to %s", route.ID, route.StartStation, route.EndStation)
	return nil, nil
}

// AdminDeleteRoute 删除路线
func AdminDeleteRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	routeId, ok := ctx.Get(RouteID).(string)
	if !ok || routeId == "" {
		log.Warn("No route ID found in context, skipping delete")
		return nil, nil
	}

	resp, err := cli.ReqDeleteRoute(routeId)
	if err != nil {
		log.Errorf("AdminDeleteRoute failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminDeleteRoute returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminDeleteRoute success: routeId=%s", routeId)
	return nil, nil
}

// generateRandomStationList 生成随机车站列表
func generateRandomStationList() []string {
	allStations := []string{
		"nanjing", "shijiazhuang", "wuxi", "shanghaihongqiao", "jiaxingnan",
		"hangzhou", "shanghai", "zhenjiang", "suzhou", "taiyuan",
		"xuzhou", "jinan", "beijing",
	}

	// 随机选择 3-6 个车站
	numStations := rand.Intn(4) + 3
	if numStations > len(allStations) {
		numStations = len(allStations)
	}

	// 随机打乱并选取
	shuffled := make([]string, len(allStations))
	copy(shuffled, allStations)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:numStations]
}

// generateRandomDistances 生成随机距离列表
func generateRandomDistances(numStations int) string {
	distances := make([]string, numStations)
	totalDistance := 0
	for i := 0; i < numStations; i++ {
		totalDistance += rand.Intn(200) + 50
		distances[i] = fmt.Sprintf("%d", totalDistance)
	}
	return strings.Join(distances, ",")
}
