package service

import (
	"testing"
)

func TestSvcImpl_CreateTravel(t *testing.T) {
	cli, _ := GetBasicClient()

	// Create Travel
	resp, err := cli.CreateTravel(&AdminTravelInfo{
		TripID: "12345",
		//LoginID             string `json:"loginId"`
		//TripID              string `json:"tripId"`
		//TrainTypeName       string `json:"trainTypeName"`
		//RouteID             string `json:"routeId"`
		//StartStationName    string `json:"startStationName"`
		//StationsName        string `json:"stationsName"`
		//TerminalStationName string `json:"terminalStationName"`
		//StartTime           string `json:"startTime"`
		//EndTime             string `json:"endTime"`
	})
	if err != nil {
		t.Errorf("CreateTravel failed: %v", err)
	}
	t.Logf("CreateTravel response: %+v", resp)

	// Update Travel
	updateResp, err := cli.UpdateTravel(&AdminTravelInfo{
		TripID: "12345",
		//LoginID             string `json:"loginId"`
		//TripID              string `json:"tripId"`
		//TrainTypeName       string `json:"trainTypeName"`
		//RouteID             string `json:"routeId"`
		//StartStationName    string `json:"startStationName"`
		//StationsName        string `json:"stationsName"`
		//TerminalStationName string `json:"terminalStationName"`
		//StartTime           string `json:"startTime"`
		//EndTime             string `json:"endTime"`
	})
	if err != nil {
		t.Errorf("UpdateTravel failed: %v", err)
	}
	t.Logf("UpdateTravel response: %+v", updateResp)

	// Delete Travel
	deleteResp, err := cli.DeleteTravel("12345")
	if err != nil {
		t.Errorf("DeleteTravel failed: %v", err)
	}
	t.Logf("DeleteTravel response: %+v", deleteResp)

	// Get All Travels
	travels, err := cli.GetAllTravels()
	if err != nil {
		t.Errorf("GetAllTravels failed: %v", err)
	}
	t.Logf("GetAllTravels response: %+v", travels)
}
