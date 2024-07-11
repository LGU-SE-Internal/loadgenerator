package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"testing"
)

func TestBasicServiceFullIntegration(t *testing.T) {
	cli, _ := GetBasicClient()
	//var basicSvc BasicService = cli

	var stationSvc StationService = cli
	stations, err := stationSvc.QueryStations()
	if err != nil {
		t.Error(err)
	}
	//t.Log(stations)

	if len(stations.Data) < 2 {
		t.Fatal("stations length should be greater than 2")
	}

	var trainSvc TrainService = cli
	trainTypes, err := trainSvc.Query()
	if err != nil {
		t.Error(err)
	}
	//t.Logf("trainTypes returns: %v", trainTypes)

	var routeSvc RouteService = cli
	// Create
	MockedID := faker.UUIDHyphenated()
	MockedStartStation := faker.GetRealAddress().City
	MockedEndStation := faker.GetRealAddress().City
	MockedStationList := fmt.Sprintf("%s,%s,%s", MockedStartStation, faker.GetRealAddress().City, MockedEndStation)
	MockedDistanceList := fmt.Sprintf("%d,%d,%d", rand.Intn(30), rand.Intn(30), rand.Intn(30))
	input := RouteInfo{
		ID:           MockedID,
		StartStation: MockedStartStation,
		EndStation:   MockedEndStation,
		StationList:  MockedStationList,
		DistanceList: MockedDistanceList,
	}
	resp, err := routeSvc.CreateAndModifyRoute(&input)
	if err != nil {
		t.Error(err)
	}
	if resp.Msg == "Already exists" {
		t.Log("Route already exists, skip")
		t.Skip()
	}
	if resp.Data.Id != input.ID {
		t.Errorf("Route ID does not match, expect %s, got %s", input.ID, resp.Data.Id)
	}
	if resp.Data.StartStation != input.StartStation {
		t.Errorf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
	}
	if resp.Data.EndStation != input.EndStation {
		t.Errorf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
	}
	if StringSliceToString(resp.Data.Stations) != ConvertCommaSeparatedToBracketed(input.StationList) {
		t.Errorf("StationList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.StationList), StringSliceToString(resp.Data.Stations))
	}
	if IntSliceToString(resp.Data.Distances) != ConvertCommaSeparatedToBracketed(input.DistanceList) {
		t.Errorf("DistanceList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.DistanceList), IntSliceToString(resp.Data.Distances))
	}
	//t.Log(resp)
	existedRoute := resp.Data

	routes, err := routeSvc.QueryAllRoutes()
	if err != nil {
		t.Error(err)
	}
	if routes.Status != 1 {
		t.Errorf("AllRoutes_By_Query.Status != 1")
	}
	found := false
	for _, route := range routes.Data {
		if route.Id == existedRoute.Id {
			found = true
		}
	}
	if !found {
		t.Errorf("Route not found by queryALL")
	}
	//t.Log(routes)

	// Mock data
	//MockedTripId := faker.UUIDHyphenated()
	MockedTripTripId := GenerateTripId()
	MockedTripTripIdType := MockedTripTripId[0]
	MockedTripTripIdNumber := MockedTripTripId[1:]
	//Input
	travelQuery := &Travel{
		Trip: Trip{
			Id: existedRoute.Id,
			TripId: TripId{
				Type:   fmt.Sprintf("%c", MockedTripTripIdType),
				Number: MockedTripTripIdNumber,
			},
			TrainTypeName:       trainTypes.Data[0].Name,
			RouteId:             existedRoute.Id,
			StartStationName:    existedRoute.StartStation,
			StationsName:        existedRoute.Stations[2], // only ok when there is exactly three stations
			TerminalStationName: existedRoute.EndStation,
			StartTime:           getRandomTime(),
			EndTime:             getRandomTime(),
		},
		StartPlace:    faker.GetRealAddress().City,
		EndPlace:      faker.GetRealAddress().City,
		DepartureTime: "",
	}

	var basicSvc BasicService = cli
	travel, err := basicSvc.QueryForTravel(travelQuery)
	if err != nil {
		t.Error(err)
	}
	if travel.Status != 1 {
		t.Log("travel.Status != 1")
	}
	//t.Log(travel)

	travels, err := basicSvc.QueryForTravels([]*Travel{travelQuery})
	if err != nil {
		t.Error(err)
	}
	if travels.Status != 1 {
		t.Log("travels.Status != 1")
	}
	//t.Log(travels)
}
