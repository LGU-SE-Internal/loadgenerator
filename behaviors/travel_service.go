package behaviors

import (
	"errors"
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
)

func QueryTripId(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.SeatService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	_ = cli

	ctx.Set(TripId, "D1345")
	return nil, nil
}

func QueryTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	tripInfo := service.TripInfo{
		StartPlace:    ctx.Get(StartStation).(string),
		EndPlace:      ctx.Get(EndStation).(string),
		DepartureTime: ctx.Get(DepartureTime).(string),
	}
	queryInfoResp, err := cli.QueryInfo(tripInfo)
	if err != nil {
		log.Errorf("QueryInfo request failed, err %s", err)
		return nil, err
	}
	if queryInfoResp.Status != 1 {
		log.Errorf("QueryInfo failed, status: %d", queryInfoResp.Status)
		return nil, err
	}

	if len(queryInfoResp.Data) == 0 {
		log.Errorf("QueryInfo response is empty")
		return nil, errors.New("QueryInfo response is empty")
	}

	randomIndex := rand.Intn(len(queryInfoResp.Data))
	ctx.Set(TripId, fmt.Sprintf("%s%s", queryInfoResp.Data[randomIndex].TripId.Type, queryInfoResp.Data[randomIndex].TripId.Number))
	//ctx.Set(TrainTypeName, queryInfoResp.Data[randomIndex].TrainTypeName)
	ctx.Set(StartStation, queryInfoResp.Data[randomIndex].StartStation)
	ctx.Set(TerminalStation, queryInfoResp.Data[randomIndex].TerminalStation)
	ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)

	return nil, nil
}

func CreateTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock para
	MockedLoginId := ctx.Get(LoginToken).(string)
	//MockedTripId := GenerateTripId()
	MockedTripId := ctx.Get(TripId).(string)
	//MockedTrainTypeName := generateTrainTypeName(MockedTripId) /*"GaoTieSeven"*/
	MockedTrainTypeName := ctx.Get(TrainTypeName).(string)
	MockedRouteID := ctx.Get(RouteID).(string)
	MockedStartStationName := ctx.Get(From).(string)
	MockedStationsName := ctx.Get(StationName).([]string)
	MockedTerminalStationName := ctx.Get(To).(string)
	MockedStartTime := ctx.Get(StartTime).(string)
	MockedEndTime := ctx.Get(EndTime).(string)

	// Mock input
	travelInfo := service.TravelInfo{
		LoginID:          MockedLoginId,
		TripID:           MockedTripId,
		TrainTypeName:    MockedTrainTypeName,
		RouteID:          MockedRouteID,
		StartStationName: MockedStartStationName,
		//StationsName:        fmt.Sprintf("%v,%v", MockedStartStationName, MockedTerminalStationName),
		StationsName:        strings.Join(MockedStationsName, ","),
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := cli.CreateTrip(&travelInfo)
	if err != nil {
		log.Errorf("CreateTrip request failed, err %s", err)
		return nil, err
	}
	if createResp.Status != 1 {
		log.Errorf("CreateTrip failed: %s", createResp.Msg)
		return nil, err
	}
	if createResp.Msg == "Already exists" {
		log.Errorf("Already exists: %s", createResp.Msg)
		return nil, err
	}
	isMatch := false
	if /*createResp.Data.Id == travelInfo.LoginID &&*/
	createResp.Data.StationsName == toLowerCaseAndRemoveSpaces(travelInfo.StationsName) &&
		createResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(travelInfo.StartStationName) &&
		createResp.Data.TerminalStationName == toLowerCaseAndRemoveSpaces(travelInfo.TerminalStationName) &&
		createResp.Data.StartTime == travelInfo.StartTime &&
		createResp.Data.EndTime == travelInfo.EndTime &&
		createResp.Data.TrainTypeName == travelInfo.TrainTypeName &&
		createResp.Data.RouteId == travelInfo.RouteID {
		isMatch = true
	}
	if !isMatch {
		log.Errorf("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
		return nil, err
	}

	/*	EndTime             string `json:"endTime"`
		Id                  string `json:"id"`
		RouteId             string `json:"routeId"`
		StartStationName    string `json:"startStationName"`
		StartTime           string `json:"startTime"`
		StationsName        string `json:"stationsName"`
		TerminalStationName string `json:"terminalStationName"`
		TrainTypeName       string `json:"trainTypeName"`
		TripId              TripId `json:"tripId"`*/

	//ctx.Set(TripId, createResp.Data.TripId)
	//ctx.Set(TrainTypeName, createResp.Data.TrainTypeName)
	//ctx.Set(StartStation, createResp.Data.StartStation)
	//ctx.Set(TerminalStation, createResp.Data.TerminalStation)
	//ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	//ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	//ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	//ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	//ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	//ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)
	return nil, nil
}
