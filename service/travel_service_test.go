package service

import (
	"github.com/go-faker/faker/v4"
	"sync"
	"testing"
)

func TestTravelService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	var travelSvc TravelService = cli

	// Query Test
	resp, err := travelSvc.QueryAll()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("Request failed, status: %d", resp.Status)
	}

	// Mock para
	MockedLoginId := faker.UUIDHyphenated()
	MockedTrainTypeName := GenerateTrainTypeName()
	MockedRouteID := faker.UUIDHyphenated()
	MockedStartStationName := faker.GetRealAddress().City
	MockedStationsName := faker.GetRealAddress().City
	MockedTerminalStationName := faker.GetRealAddress().City
	MockedStartTime := getRandomTime()
	MockedEndTime := getRandomTime()
	MockedTripId := GenerateTripId()

	// Mock input
	travelInfo := TravelInfo{
		LoginID:          MockedLoginId,
		TripID:           MockedTripId,
		TrainTypeName:    MockedTrainTypeName,
		RouteID:          MockedRouteID,
		StartStationName: MockedStartStationName,
		//StationsName:        fmt.Sprintf("%v,%v", MockedStartStationName, MockedTerminalStationName),
		StationsName:        MockedStationsName,
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := travelSvc.CreateTrip(&travelInfo)
	if err != nil {
		t.Errorf("CreateTrip request failed, err %s", err)
	}
	if createResp.Status != 1 {
		t.Errorf("CreateTrip failed: %s", createResp.Msg)
	}
	//if createResp.Msg != "Already exists" {
	//	t.Logf("Already exists: %s", createResp.Msg)
	//	t.Skip()
	//}
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
		t.Errorf("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
	}
	existedTravel := createResp.Data

	// Query all
	allTravelInfos, err := cli.QueryAll()
	if err != nil {
		t.Errorf("QueryAll request failed, err %s", err)
	}
	if len(allTravelInfos.Data) == 0 {
		t.Errorf("QueryAll returned no results")
	}
	if allTravelInfos.Status != 1 {
		t.Errorf("QueryAll failed, status: %d", allTravelInfos.Status)
	}
	found := false
	for _, travel := range allTravelInfos.Data {
		if /*travel.Id == existedTravel.Id &&*/
		travel.RouteId == existedTravel.RouteId &&
			travel.StartTime == existedTravel.StartTime &&
			travel.EndTime == existedTravel.EndTime &&
			travel.TrainTypeName == existedTravel.TrainTypeName &&
			travel.StartStationName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&
			/*travel.StationsName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&*/
			travel.TerminalStationName == toLowerCaseAndRemoveSpaces(existedTravel.TerminalStationName) {
			found = true
		}
	}
	if !found {
		t.Errorf("Cannot find existed travel info: %v", existedTravel)
	}

	// QueryAll - Admin
	adminQueryAllResp, err := cli.AdminQueryAll()
	if err != nil {
		t.Errorf("AdminQueryAll request failed, err %s", err)
	}
	if adminQueryAllResp.Status != 1 {
		t.Errorf("AdminQueryAll failed, status: %d", adminQueryAllResp.Status)
	}
	found2 := false
	for _, travel := range adminQueryAllResp.Data {
		if travel.Trip.Id == existedTravel.Id &&
			travel.Trip.RouteId == existedTravel.RouteId &&
			travel.Trip.StartTime == existedTravel.StartTime &&
			travel.Trip.EndTime == existedTravel.EndTime &&
			travel.Trip.TrainTypeName == existedTravel.TrainTypeName &&
			travel.Trip.StartStationName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&
			/*travel.Trip.StationsName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&*/
			travel.Trip.TerminalStationName == toLowerCaseAndRemoveSpaces(existedTravel.TerminalStationName) {
			found2 = true
		}
	}
	if !found2 {
		t.Errorf("Cannot find existed travel info: %v", existedTravel)
	}

	// Test Update
	// Update the para
	MockedStartStationName = faker.GetRealAddress().City
	MockedStationsName = faker.GetRealAddress().City
	MockedTerminalStationName = faker.GetRealAddress().City
	updateTravelInfo := &TravelInfo{
		LoginID:             MockedLoginId,
		TripID:              MockedTripId,
		TrainTypeName:       MockedTrainTypeName,
		RouteID:             MockedRouteID,
		StartStationName:    MockedStartStationName,
		StationsName:        MockedStationsName,
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
	isMatch1 := false
	if updateResp.Data.Id == existedTravel.Id &&
		/*updateResp.Data.StationsName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&*/
		updateResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(existedTravel.StartStationName) &&
		updateResp.Data.TerminalStationName == toLowerCaseAndRemoveSpaces(existedTravel.TerminalStationName) &&
		updateResp.Data.StartTime == existedTravel.StartTime &&
		updateResp.Data.EndTime == existedTravel.EndTime &&
		updateResp.Data.TrainTypeName == existedTravel.TrainTypeName &&
		updateResp.Data.RouteId == existedTravel.RouteId {
		isMatch1 = true
	}
	if !isMatch1 {
		t.Errorf("Expect: %v, get: %v", travelInfo, updateResp.Data)
	}

	// Query all UpdatedInfo
	allUpdatedTravelInfos, err := cli.QueryAll()
	if err != nil {
		t.Errorf("QueryAll request failed, err %s", err)
	}
	if len(allUpdatedTravelInfos.Data) == 0 {
		t.Errorf("QueryAll returned no results")
	}
	if allUpdatedTravelInfos.Status != 1 {
		t.Errorf("QueryAll failed, status: %d", allUpdatedTravelInfos.Status)
	}
	found1 := false
	for _, travel := range allUpdatedTravelInfos.Data {
		if travel.Id == existedTravel.Id &&
			travel.RouteId == existedTravel.RouteId &&
			travel.StartTime == existedTravel.StartTime &&
			travel.EndTime == existedTravel.EndTime &&
			travel.TrainTypeName == existedTravel.TrainTypeName &&
			travel.StartStationName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&
			/*travel.StationsName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&*/
			travel.TerminalStationName == toLowerCaseAndRemoveSpaces(existedTravel.TerminalStationName) {
			found1 = true
		}
	}
	if !found1 {
		t.Errorf("Cannot find existed travel info: %v", existedTravel)
	}

	// Test Retrieve by ID
	retrieveResp, err := cli.Retrieve(MockedLoginId)
	if err != nil {
		t.Errorf("Retrieve request failed, err %s", err)
	}
	if retrieveResp.Status != 1 {
		t.Errorf("Retrieve failed: %s", retrieveResp.Msg)
	}

	// Test GetTrainTypeByTripId
	trainTypeResp, err := cli.GetTrainTypeByTripId(MockedTripId)
	if err != nil {
		t.Errorf("GetTrainTypeByTripId request failed, err %s", err)
	}
	if trainTypeResp.Status != 1 {
		t.Errorf("GetTrainTypeByTripId failed, status: %d", trainTypeResp.Status)
	}

	// Test GetRouteByTripId
	routeResp, err := cli.GetRouteByTripId(MockedTripId)
	if err != nil {
		t.Errorf("GetRouteByTripId request failed, err %s", err)
	}
	if routeResp == nil {
		t.Errorf("GetRouteByTripId returned no result")
	}

	// Test GetTripsByRouteId
	routeIds := []string{MockedRouteID}
	tripsByRouteResp, err := cli.GetTripsByRouteId(routeIds)
	if err != nil {
		t.Errorf("GetTripsByRouteId request failed, err %s", err)
	}
	if len(tripsByRouteResp.Data) == 0 {
		t.Errorf("GetTripsByRouteId returned no results")
	}

	// Test QueryInfo
	tripInfo := TripInfo{
		StartPlace:    MockedStationsName,
		EndPlace:      MockedTerminalStationName,
		DepartureTime: MockedStartTime,
	}
	queryInfoResp, err := cli.QueryInfo(tripInfo)
	if err != nil {
		t.Errorf("QueryInfo request failed, err %s", err)
	}

	// Test QueryInfoInParallel
	queryInfoInParallelResp, err := cli.QueryInfoInParallel(tripInfo)
	if err != nil {
		t.Errorf("QueryInfoInParallel request failed, err %s", err)
	}
	if queryInfoInParallelResp.Status != 1 {
		t.Errorf("QueryInfoInParallel failed, status: %d", queryInfoResp.Status)
	}

	// Test GetTripAllDetailInfo
	//tripAllDetailResp, err := cli.GetTripAllDetailInfo(GetTripDetailReq{
	//	From:       "suzhou",
	//	To:         "taiyuan",
	//	TravelDate: "",
	//	TripId:     "G1234",
	//})
	tripAllDetailResp, err := cli.GetTripAllDetailInfo(GetTripDetailReq{
		From:       MockedStationsName,
		To:         MockedTerminalStationName,
		TravelDate: MockedStartTime,
		TripId:     MockedTripId,
	})
	if err != nil {
		t.Errorf("GetTripAllDetailInfo request failed, err %s", err)
	}
	if tripAllDetailResp.Status != 1 {
		t.Errorf("GetTripAllDetailInfo failed, status: %d", tripAllDetailResp.Status)
	}

	// Test Delete
	deleteResp, err := cli.DeleteTrip(MockedTripId)
	if err != nil {
		t.Errorf("DeleteTrip request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("DeleteTrip failed: %s", deleteResp.Msg)
	}
	t.Logf("DeleteTrip return: %s", deleteResp.Msg)

}

func TestTravelServiceQueryAll_InfiniteLoop_ForTesting(t *testing.T) {
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
