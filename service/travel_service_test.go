package service

import (
	"github.com/go-faker/faker/v4"
	"sync"
	"testing"
)

func TestTravelServiceQueryAll(t *testing.T) {
	var wg sync.WaitGroup
	//numIterations := 100

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			cli, _ := GetAdminClient()
			for {
				// Query Test
				_, err := cli.QueryAll()
				if err != nil {
					t.Errorf("Request failed, err %s", err)
					return
				}
				// Process the response if needed
			}
		}()

		wg.Wait()
	}

}

func TestTravelService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()

	// Query Test
	resp, err := cli.QueryAll()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	t.Logf("QueryAll return: %+v", resp)

	MockedLoginId := faker.UUIDHyphenated()
	MockedTrainTypeName := faker.Word()
	MockedRouteID := faker.UUIDHyphenated()
	MockedStartStationName := "Shenzhen Bei"
	MockedTerminalStationName := "California Airport"
	MockedStartTime := faker.Date()
	MockedEndTime := faker.Date()
	TripId := "G777"

	travelInfo := &TravelInfo{
		LoginID:             MockedLoginId,
		TripID:              TripId,
		TrainTypeName:       MockedTrainTypeName,
		RouteID:             MockedRouteID,
		StartStationName:    MockedStartStationName,
		StationsName:        "Shenzhen Bei, California Airport",
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := cli.CreateTrip(travelInfo)
	if err != nil {
		t.Errorf("CreateTrip request failed, err %s", err)
	}
	if createResp.Status != 1 {
		t.Errorf("CreateTrip failed: %s", createResp.Msg)
	}

	// Query all
	allTravelInfos, err := cli.QueryAll()
	if err != nil {
		t.Errorf("QueryAll request failed, err %s", err)
	}
	if len(allTravelInfos.Data) == 0 {
		t.Errorf("QueryAll returned no results")
	}

	// Test Update
	updateTravelInfo := &TravelInfo{
		LoginID:             MockedLoginId,
		TripID:              TripId,
		TrainTypeName:       MockedTrainTypeName,
		RouteID:             MockedRouteID,
		StartStationName:    MockedStartStationName,
		StationsName:        "Shenzhen Bei,Futian",
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}
	updateResp, err := cli.UpdateTrip(updateTravelInfo)
	if err != nil {
		t.Errorf("UpdateTrip request failed, err %s", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("UpdateTrip failed: %s", updateResp.Msg)
	}

	// Test Delete
	var deleteID string
	if len(allTravelInfos.Data) > 0 {
		deleteID = allTravelInfos.Data[len(allTravelInfos.Data)-1].TripId.Type + allTravelInfos.Data[len(allTravelInfos.Data)-1].TripId.Number
	} else {
		t.Errorf("QueryAll returned empty data")
	}
	deleteResp, err := cli.DeleteTrip(deleteID)
	if err != nil {
		t.Errorf("DeleteTrip request failed, err %s", err)
	}
	t.Logf("DeleteTrip return: %s", deleteResp.Msg)

	// Test Retrieve by ID
	retrieveResp, err := cli.Retrieve(TripId)
	if err != nil {
		t.Errorf("Retrieve request failed, err %s", err)
	}
	if retrieveResp == nil {
		t.Errorf("Retrieve returned no result")
	}

	// Test GetTrainTypeByTripId
	trainTypeResp, err := cli.GetTrainTypeByTripId(TripId)
	if err != nil {
		t.Errorf("GetTrainTypeByTripId request failed, err %s", err)
	}
	if trainTypeResp == nil {
		t.Errorf("GetTrainTypeByTripId returned no result")
	}

	// Test GetRouteByTripId
	routeResp, err := cli.GetRouteByTripId(TripId)
	if err != nil {
		t.Errorf("GetRouteByTripId request failed, err %s", err)
	}
	if routeResp == nil {
		t.Errorf("GetRouteByTripId returned no result")
	}

	// Test GetTripsByRouteId
	routeIds := []string{faker.UUIDHyphenated(), faker.UUIDHyphenated()}
	tripsByRouteResp, err := cli.GetTripsByRouteId(routeIds)
	if err != nil {
		t.Errorf("GetTripsByRouteId request failed, err %s", err)
	}
	if len(tripsByRouteResp.Data) == 0 {
		t.Errorf("GetTripsByRouteId returned no results")
	}

	// Test QueryInfo
	tripInfo := TripInfo{
		StartPlace:    "PlaceA",
		EndPlace:      "PlaceB",
		DepartureTime: "2023-01-01",
	}
	queryInfoResp, err := cli.QueryInfo(tripInfo)
	if err != nil {
		t.Errorf("QueryInfo request failed, err %s", err)
	}
	t.Logf("QueryInfo returned results: %v", queryInfoResp)

	// Test QueryInfoInParallel
	queryInfoInParallelResp, err := cli.QueryInfoInParallel(tripInfo)
	if err != nil {
		t.Errorf("QueryInfoInParallel request failed, err %s", err)
	}
	t.Logf("QueryInfoInParallel returns: %v", queryInfoInParallelResp)

	// Test GetTripAllDetailInfo
	tripAllDetailResp, err := cli.GetTripAllDetailInfo(GetTripDetailReq{
		From:       "suzhou",
		To:         "taiyuan",
		TravelDate: "",
		TripId:     "G1234",
	})
	if err != nil {
		t.Errorf("GetTripAllDetailInfo request failed, err %s", err)
	}
	t.Logf("GetTripAllDetailInfo returns: %v", tripAllDetailResp)

	// Test AdminQueryAll
	adminQueryAllResp, err := cli.AdminQueryAll()
	if err != nil {
		t.Errorf("AdminQueryAll request failed, err %s", err)
	}
	t.Logf("AdminQueryAll returns: %v", adminQueryAllResp)

}

func TestGetTripAllDetailInfo(t *testing.T) {
	cli, _ := GetAdminClient()

	resp, err := cli.QueryAll()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	t.Logf("QueryAll return: %+v", resp)

	res, err := cli.GetTripAllDetailInfo(GetTripDetailReq{
		From:       "shanghai",
		To:         "taiyuan",
		TravelDate: "2025-05-04 09:00:00",
		TripId:     "G1234",
	})
	if err != nil {
		t.Errorf("GetTripAllDetailInfo request failed, err %s", err)
	}
	t.Log(res)
}
