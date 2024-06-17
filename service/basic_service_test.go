package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient() // Assuming GetBasicClient is implemented elsewhere

	// Mock data for Travel
	MockedTypeName := faker.Word()
	MockedIndex := 1
	MockedTripIDName := faker.Word()
	MockedTrainTypeName := faker.Word()
	MockedRouteID := faker.UUIDHyphenated()
	MockedStartStationName := "Shenzhen Bei"
	MockedTerminalStationName := "California Airport"
	MockedStartTime := faker.Date()
	MockedEndTime := faker.Date()
	MockedDepartureTime := faker.Date()
	travel := &Travel{
		Trip: Trip{
			ID: "1",
			TripID: TripId{
				Type:   Type{Name: MockedTypeName, Index: MockedIndex},
				Number: MockedTripIDName,
			},
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        "Shenzhen Bei,California Airport",
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		},
		StartPlace:    "Shenzhen",
		EndPlace:      "California",
		DepartureTime: MockedDepartureTime,
	}

	// Test QueryForTravel
	queryTravelResp, err := cli.QueryForTravel(travel)
	if err != nil {
		t.Errorf("QueryForTravel request failed, err %s", err)
	}
	t.Logf("QueryForTravel returned results: %v", queryTravelResp)

	// Mock data for Travels
	travels := []Travel{*travel}

	// Test QueryForTravels
	queryTravelsResp, err := cli.QueryForTravels(travels)
	if err != nil {
		t.Errorf("QueryForTravels request failed, err %s", err)
	}
	t.Logf("QueryForTravels returned results: %v", queryTravelsResp)

	// Mock data for Station Name
	stationName := faker.Word()

	// Test QueryForStationId
	queryStationIdResp, err := cli.QueryForStationId(stationName)
	if err != nil {
		t.Errorf("QueryForStationId request failed, err %s", err)
	}
	t.Logf("QueryForStationId returned results: %v", queryStationIdResp)
}
