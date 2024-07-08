package service

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestBasicServiceFullIntegration(t *testing.T) {
	cli, _ := GetBasicClient()

	var stationSvc StationService = cli
	stations, err := stationSvc.QueryStations()
	if err != nil {
		t.Error(err)
	}
	t.Log(stations)

	if len(stations.Data) < 2 {
		t.Fatal("stations length should be greater than 2")
	}

	var trainSvc TrainService = cli
	trainTypes, err := trainSvc.Query()
	if err != nil {
		t.Error(err)
	}
	t.Logf("trainTypes returns: %v", trainTypes)

	var routeSvc RouteService = cli
	MockedID := faker.UUIDHyphenated()
	input := &RouteInfo{
		ID:           MockedID,
		StartStation: "Shenzhen Bei",
		EndStation:   "Jiulong Xi",
		StationList:  "Shenzhen Bei,Shkou,Jiulong Xi",
		DistanceList: "77,66,55",
	}
	resp, err := routeSvc.CreateAndModifyRoute(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

	routes, err := routeSvc.QueryAllRoutes()
	if err != nil {
		t.Error(err)
	}
	t.Log(routes)

	travelQuery := &Travel{
		Trip: Trip{
			Id: "6284bf46-0f0a-481d-a221-bb3794b00585",
			TripId: TripId{
				Type:   "G",
				Number: "985",
			},
			TrainTypeName:       trainTypes.Data[0].Name,
			RouteId:             routes.Data[0].Id,
			StartStationName:    "",
			StationsName:        "",
			TerminalStationName: "",
			StartTime:           "",
			EndTime:             "",
		},
		StartPlace:    "Shenzhen Bei",
		EndPlace:      "Jiulong Xi",
		DepartureTime: "",
	}
	travel, err := cli.QueryForTravel(travelQuery)
	if err != nil {
		t.Error(err)
	}
	t.Log(travel)

	travels, err := cli.QueryForTravels([]*Travel{travelQuery})
	if err != nil {
		t.Error(err)
	}
	t.Log(travels)
}
