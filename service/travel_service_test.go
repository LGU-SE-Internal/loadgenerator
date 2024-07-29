package service

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"testing"
)

func TestTravelService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient() // The loginResult below should also be the corresponding one! Or -> Forbidden.
	var travelSvc TravelService = cli

	/*	loginResult, err := cli.ReqUserLogin(&UserLoginInfoReq{ // Basic
		Password:         "111111",
		UserName:         "fdse_microservice",
		VerificationCode: "123",
	})*/
	loginResult, err := cli.ReqUserLogin(&UserLoginInfoReq{ // Admin
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Query Test
	resp, err := travelSvc.QueryAllTrip()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("Request failed, status: %d", resp.Status)
	}

	var routeSvc RouteService = cli
	AllRoutesByQuery, err2 := routeSvc.QueryAllRoutes()
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if AllRoutesByQuery.Status != 1 {
		t.Errorf("AllRoutes_By_Query.Status != 1")
	}

	routeRandomIndex := rand.Intn(len(AllRoutesByQuery.Data))
	randomRoute := AllRoutesByQuery.Data[routeRandomIndex]

	// Mock para
	MockedLoginId := loginResult.Data.Token
	MockedTrainTypeName := GenerateTrainTypeName() /*"GaoTieSeven"*/
	MockedRouteID := randomRoute.Id
	MockedStartStationName := randomRoute.StartStation
	MockedStationsName := /*strings.Join(AllRoutesByQuery.Data[0].Stations, ",")*/ getMiddleElements(strings.Join(randomRoute.Stations, ","))
	MockedTerminalStationName := randomRoute.EndStation
	MockedStartTime := getRandomTime()
	MockedEndTime := getRandomTime(WithStartTime(MockedStartTime))
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
	if createResp.Msg == "Already exists" {
		t.Logf("Already exists: %s", createResp.Msg)
		t.Skip()
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
		t.Errorf("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
	}
	existedTravel := createResp.Data

	// Query all
	allTravelInfos, err := travelSvc.QueryAllTrip()
	if err != nil {
		t.Errorf("QueryAllTrip request failed, err %s", err)
	}
	if len(allTravelInfos.Data) == 0 {
		t.Errorf("QueryAllTrip returned no results")
	}
	if allTravelInfos.Status != 1 {
		t.Errorf("QueryAllTrip failed, status: %d", allTravelInfos.Status)
	}
	found := false
	for _, travel := range allTravelInfos.Data {
		if travel.Id == existedTravel.Id {
			found = true
		}
	}
	if !found {
		t.Errorf("Cannot find existed travel info: %v", existedTravel)
	}

	// QueryAllTrip - Admin
	adminQueryAllResp, err := travelSvc.AdminQueryAll()
	if err != nil {
		t.Errorf("AdminQueryAll request failed, err %s", err)
	}
	if adminQueryAllResp.Status != 1 {
		t.Errorf("AdminQueryAll failed, status: %d", adminQueryAllResp.Status)
	}
	found2 := false
	for _, travel := range adminQueryAllResp.Data {
		if travel.Trip.Id == existedTravel.Id {
			found2 = true
		}
	}
	if !found2 {
		t.Errorf("Cannot find existed travel info: %v", existedTravel)
	}

	// Test Update
	// Query for the Stations
	// Station Service
	var stationSvc StationService = cli

	QueryAllStations, err := stationSvc.QueryStations()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if len(resp.Data) == 0 {
		t.Errorf("No stations found")
	}
	if QueryAllStations.Status != 1 {
		t.Errorf("Status should be 1, but is %d", resp.Status)
	}

	// Generate a random index within the range of Data list length
	randomIndex := rand.Intn(len(QueryAllStations.Data))
	// Access the Name field using the random index
	randomStationName := QueryAllStations.Data[randomIndex].Name

	// Update the para
	MockedModifiedTerminalStationName := randomStationName
	updateTravelInfo := &TravelInfo{
		LoginID:             MockedLoginId,
		TripID:              fmt.Sprintf("%s%s", existedTravel.TripId.Type, existedTravel.TripId.Number),
		TrainTypeName:       existedTravel.TrainTypeName,
		RouteID:             existedTravel.RouteId,
		StartStationName:    existedTravel.StartStationName,
		StationsName:        existedTravel.StationsName,
		TerminalStationName: MockedModifiedTerminalStationName,
		StartTime:           existedTravel.StartTime,
		EndTime:             existedTravel.EndTime,
	}
	updateResp, err := travelSvc.UpdateTrip(updateTravelInfo)
	if err != nil {
		t.Errorf("UpdateTrip request failed, err %s", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("UpdateTrip failed: %s", updateResp.Msg)
	}
	isMatch1 := false
	if updateResp.Data.Id == existedTravel.Id &&
		/*updateResp.Data.StationsName == toLowerCaseAndRemoveSpaces(existedTravel.StationsName) &&*/
		//updateResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(existedTravel.StartStationName) &&
		updateResp.Data.StartStationName == existedTravel.StartStationName &&
		updateResp.Data.StartTime == existedTravel.StartTime &&
		updateResp.Data.EndTime == existedTravel.EndTime &&
		updateResp.Data.TrainTypeName == existedTravel.TrainTypeName &&
		updateResp.Data.RouteId == existedTravel.RouteId {
		isMatch1 = true
	}
	if !isMatch1 {
		t.Errorf("Expect: %v, get: %v", travelInfo, updateResp.Data)
	}
	updatedTravel := updateResp.Data

	// Query all UpdatedInfo
	allUpdatedTravelInfos, err := travelSvc.QueryAllTrip()
	if err != nil {
		t.Errorf("QueryAllTrip request failed, err %s", err)
	}
	if len(allUpdatedTravelInfos.Data) == 0 {
		t.Errorf("QueryAllTrip returned no results")
	}
	if allUpdatedTravelInfos.Status != 1 {
		t.Errorf("QueryAllTrip failed, status: %d", allUpdatedTravelInfos.Status)
	}
	found1 := false
	for _, travel := range allUpdatedTravelInfos.Data {
		if travel.Id == updatedTravel.Id &&
			travel.RouteId == updatedTravel.RouteId &&
			travel.StartTime == updatedTravel.StartTime &&
			travel.EndTime == updatedTravel.EndTime &&
			travel.TrainTypeName == updatedTravel.TrainTypeName &&
			travel.StationsName == updatedTravel.StationsName {
			found1 = true
		}
	}
	if !found1 {
		t.Errorf("Cannot find existed travel info: %v", existedTravel)
	}

	// Test GetTrainTypeByTripId - K, Z
	trainTypeResp, err := travelSvc.GetTrainTypeByTripId(fmt.Sprintf("%s%s", updatedTravel.TripId.Type, updatedTravel.TripId.Number))
	if err != nil {
		t.Errorf("GetTrainTypeByTripId request failed, err %s", err)
	}
	if trainTypeResp.Status != 1 {
		t.Errorf("GetTrainTypeByTripId failed, status: %d", trainTypeResp.Status)
	}

	// Test GetRouteByTripId
	routeResp, err := travelSvc.GetRouteByTripId(fmt.Sprintf("%s%s", updatedTravel.TripId.Type, updatedTravel.TripId.Number))
	if err != nil {
		t.Errorf("GetRouteByTripId request failed, err %s", err)
	}
	if routeResp == nil {
		t.Errorf("GetRouteByTripId returned no result")
	}
	if routeResp.Status != 1 {
		t.Errorf("GetRouteByTripId failed, status: %d", routeResp.Status)
	}

	// Test GetTripsByRouteId
	routeIds := []string{updatedTravel.RouteId}
	tripsByRouteResp, err := travelSvc.GetTripsByRouteId(routeIds)
	if err != nil {
		t.Errorf("GetTripsByRouteId request failed, err %s", err)
	}
	if len(tripsByRouteResp.Data) == 0 {
		t.Errorf("GetTripsByRouteId returned no results")
	}
	if tripsByRouteResp.Status != 1 {
		t.Errorf("GetTripsByRouteId failed, status: %d", tripsByRouteResp.Status)
	}

	// Test QueryInfo
	tripInfo := TripInfo{
		StartPlace:    updatedTravel.StartStationName,
		EndPlace:      updatedTravel.TerminalStationName,
		DepartureTime: updatedTravel.StartTime,
	}
	queryInfoResp, err := travelSvc.QueryInfo(tripInfo)
	if err != nil {
		t.Errorf("QueryInfo request failed, err %s", err)
	}
	if queryInfoResp.Status != 1 {
		t.Errorf("QueryInfo failed, status: %d", queryInfoResp.Status)
	}

	// Test QueryInfoInParallel
	queryInfoInParallelResp, err := travelSvc.QueryInfoInParallel(tripInfo)
	if err != nil {
		t.Errorf("QueryInfoInParallel request failed, err %s", err)
	}
	if queryInfoInParallelResp.Status != 1 {
		t.Errorf("QueryInfoInParallel failed, status: %d", queryInfoInParallelResp.Status)
	}

	// Test GetTripAllDetailInfo
	//tripAllDetailResp, err := cli.GetTripAllDetailInfo(GetTripDetailReq{
	//	From:       "suzhou",
	//	To:         "taiyuan",
	//	TravelDate: "",
	//	TripId:     "G1234",
	//})
	tripAllDetailResp, err := travelSvc.GetTripAllDetailInfo(GetTripDetailReq{
		From:       updatedTravel.StartStationName,
		To:         updatedTravel.TerminalStationName,
		TravelDate: updatedTravel.StartTime,
		TripId:     fmt.Sprintf("%s%s", updatedTravel.TripId.Type, updatedTravel.TripId.Number),
	})
	if err != nil {
		t.Errorf("GetTripAllDetailInfo request failed, err %s", err)
	}
	if tripAllDetailResp.Status != 1 {
		t.Errorf("GetTripAllDetailInfo failed, status: %d", tripAllDetailResp.Status)
	}

	// Test Delete
	deleteResp, err := travelSvc.DeleteTrip(MockedTripId)
	if err != nil {
		t.Errorf("DeleteTrip request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("DeleteTrip failed: %s", deleteResp.Msg)
	}
	t.Logf("DeleteTrip successfully return: %s", deleteResp.Msg)

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
				_, err := cli.QueryAllTrip()
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

	resp, err := cli.QueryAllTrip()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	t.Logf("QueryAllTrip return: %+v", resp)

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
