package behaviors

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// Admin Travel Behavior Chain - 管理行程的增删改查操作
var AdminTravelChain *Chain

func init() {
	AdminTravelChain = NewChain(
		NewFuncNode(LoginAdmin, "LoginAdmin"),
		NewFuncNode(AdminQueryAllTravels, "AdminQueryAllTravels"),
	)

	// 添加后续的行程管理操作链
	AdminTravelChain.AddNextChain(NewChain(
		NewFuncNode(AdminCreateTravel, "AdminCreateTravel"),
		NewFuncNode(AdminUpdateTravel, "AdminUpdateTravel"),
	), 0.5)

	AdminTravelChain.AddNextChain(NewChain(
		NewFuncNode(AdminDeleteTravel, "AdminDeleteTravel"),
	), 0.5)
}

// AdminQueryAllTravels 查询所有行程
func AdminQueryAllTravels(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllTravels()
	if err != nil {
		log.Errorf("AdminQueryAllTravels failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminQueryAllTravels returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	if len(resp.Data) > 0 {
		randomIndex := rand.Intn(len(resp.Data))
		ctx.Set(TripID, resp.Data[randomIndex].TripID)
		ctx.Set(TrainTypeName, resp.Data[randomIndex].TrainTypeName)
		ctx.Set(RouteID, resp.Data[randomIndex].RouteID)
		ctx.Set(StartStation, resp.Data[randomIndex].StartStationName)
		ctx.Set(EndStation, resp.Data[randomIndex].TerminalStationName)
		ctx.Set(StartTime, resp.Data[randomIndex].StartTime)
		ctx.Set(EndTime, resp.Data[randomIndex].EndTime)
	}

	log.Infof("AdminQueryAllTravels returned %d travels", len(resp.Data))
	return nil, nil
}

// AdminCreateTravel 创建行程
func AdminCreateTravel(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	stations := generateRandomStationList()

	travel := &service.AdminTravelInfo{
		LoginID:             ctx.Get(LoginToken).(string),
		TripID:              GenerateTripId(),
		TrainTypeName:       GenerateTrainTypeName(),
		RouteID:             uuid.New().String(),
		StartStationName:    stations[0],
		StationsName:        strings.Join(stations, ","),
		TerminalStationName: stations[len(stations)-1],
		StartTime:           getRandomTime(),
		EndTime:             getRandomTime(WithStartTime(getRandomTime())),
	}

	resp, err := cli.CreateTravel(travel)
	if err != nil {
		log.Errorf("AdminCreateTravel failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminCreateTravel returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(TripID, travel.TripID)
	ctx.Set(TrainTypeName, travel.TrainTypeName)
	ctx.Set(RouteID, travel.RouteID)
	ctx.Set(StartStation, travel.StartStationName)
	ctx.Set(EndStation, travel.TerminalStationName)
	ctx.Set(StationName, stations)

	log.Infof("AdminCreateTravel success: tripId=%s, from %s to %s", travel.TripID, travel.StartStationName, travel.TerminalStationName)
	return nil, nil
}

// AdminUpdateTravel 更新行程
func AdminUpdateTravel(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	tripId, ok := ctx.Get(TripID).(string)
	if !ok || tripId == "" {
		log.Warn("No trip ID found in context, skipping update")
		return nil, nil
	}

	stations := generateRandomStationList()

	travel := &service.AdminTravelInfo{
		LoginID:             ctx.Get(LoginToken).(string),
		TripID:              tripId,
		TrainTypeName:       GenerateTrainTypeName(),
		RouteID:             uuid.New().String(),
		StartStationName:    stations[0],
		StationsName:        strings.Join(stations, ","),
		TerminalStationName: stations[len(stations)-1],
		StartTime:           getRandomTime(),
		EndTime:             getRandomTime(WithStartTime(getRandomTime())),
	}

	resp, err := cli.UpdateTravel(travel)
	if err != nil {
		log.Errorf("AdminUpdateTravel failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminUpdateTravel returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminUpdateTravel success: tripId=%s", tripId)
	return nil, nil
}

// AdminDeleteTravel 删除行程
func AdminDeleteTravel(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	tripId, ok := ctx.Get(TripID).(string)
	if !ok || tripId == "" {
		log.Warn("No trip ID found in context, skipping delete")
		return nil, nil
	}

	resp, err := cli.DeleteTravel(tripId)
	if err != nil {
		log.Errorf("AdminDeleteTravel failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminDeleteTravel returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminDeleteTravel success: tripId=%s", tripId)
	return nil, nil
}
