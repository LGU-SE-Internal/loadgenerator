package service

//func TestTravelService_FullIntegration(t *testing.T) {
//	cli, _ := GetBasicClient()
//
//	// Query Test
//	resp, err := cli.QueryAll()
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if len(resp.Data) > 0 {
//		t.Errorf("QueryAll returned no results")
//	}
//
//	// Mock data
//	//MockedTypeName := faker.Word()
//	MockedTripID := faker.UUIDHyphenated()
//	//MockedIndex := 1
//	//MockedTripIDName := faker.Word()
//	MockedTrainTypeName := faker.Word()
//	MockedRouteID := faker.UUIDHyphenated()
//	MockedStartStationName := "Shenzhen Bei"
//	MockedTerminalStationName := "California Airport"
//	MockedStartTime := faker.Date()
//	MockedEndTime := faker.Date()
//
//	travelInfo := &TravelInfo{
//		LoginID:             "1",
//		TripID:              MockedTripID,
//		TrainTypeName:       MockedTrainTypeName,
//		RouteID:             MockedRouteID,
//		StartStationName:    MockedStartStationName,
//		StationsName:        "Shenzhen Bei,California Airport",
//		TerminalStationName: MockedTerminalStationName,
//		StartTime:           MockedStartTime,
//		EndTime:             MockedEndTime,
//	}
//
//	// Create Test
//	createResp, err := cli.CreateTrip(travelInfo)
//	if err != nil {
//		t.Errorf("CreateTrip request failed, err %s", err)
//	}
//	if createResp.Status != 1 {
//		t.Errorf("CreateTrip failed: %s", createResp.Msg)
//	}
//
//	// Query all
//	allTravelInfos, err := cli.QueryAll()
//	if err != nil {
//		t.Errorf("QueryAll request failed, err %s", err)
//	}
//	if len(allTravelInfos.Data) == 0 {
//		t.Errorf("QueryAll returned no results")
//	}
//
//	var getId string
//	if len(allTravelInfos.Data) > 0 {
//		getId = allTravelInfos.Data[0].DestStation
//	}
//
//	// Test Update
//	updateTravelInfo := &TravelInfo{
//		LoginID:             "1",
//		TripID:              MockedTripID,
//		TrainTypeName:       MockedTrainTypeName,
//		RouteID:             MockedRouteID,
//		StartStationName:    MockedStartStationName,
//		StationsName:        "Shenzhen Bei,Futian",
//		TerminalStationName: MockedTerminalStationName,
//		StartTime:           MockedStartTime,
//		EndTime:             MockedEndTime,
//	}
//	updateResp, err := cli.UpdateTrip(updateTravelInfo)
//	if err != nil {
//		t.Errorf("UpdateTrip request failed, err %s", err)
//	}
//	if updateResp.Status != 1 {
//		t.Errorf("UpdateTrip failed: %s", updateResp.Msg)
//	}
//
//	// Test Delete
//	var deleteID string
//	if len(allTravelInfos.Data) > 0 {
//		deleteID = allTravelInfos.Data[len(allTravelInfos.Data)-1].DestStation
//	} else {
//		t.Errorf("QueryAll returned empty data")
//	}
//	deleteResp, err := cli.DeleteTrip(deleteID)
//	if err != nil {
//		t.Errorf("DeleteTrip request failed, err %s", err)
//	}
//	if deleteResp.Status != 1 {
//		t.Errorf("DeleteTrip failed: %s", deleteResp.Msg)
//	}
//
//	// Test Retrieve by ID
//	retrieveResp, err := cli.Retrieve(getId)
//	if err != nil {
//		t.Errorf("Retrieve request failed, err %s", err)
//	}
//	if retrieveResp == nil {
//		t.Errorf("Retrieve returned no result")
//	}
//
//	// Test GetTrainTypeByTripId
//	trainTypeResp, err := cli.GetTrainTypeByTripId(getId)
//	if err != nil {
//		t.Errorf("GetTrainTypeByTripId request failed, err %s", err)
//	}
//	if trainTypeResp == nil {
//		t.Errorf("GetTrainTypeByTripId returned no result")
//	}
//
//	// Test GetRouteByTripId
//	routeResp, err := cli.GetRouteByTripId(getId)
//	if err != nil {
//		t.Errorf("GetRouteByTripId request failed, err %s", err)
//	}
//	if routeResp == nil {
//		t.Errorf("GetRouteByTripId returned no result")
//	}
//
//	// Test GetTripsByRouteId
//	routeIds := []string{faker.UUIDHyphenated(), faker.UUIDHyphenated()}
//	tripsByRouteResp, err := cli.GetTripsByRouteId(routeIds)
//	if err != nil {
//		t.Errorf("GetTripsByRouteId request failed, err %s", err)
//	}
//	if len(tripsByRouteResp.Data) == 0 {
//		t.Errorf("GetTripsByRouteId returned no results")
//	}
//
//	// Test QueryInfo
//	tripInfo := TripInfo{
//		StartPlace:    "PlaceA",
//		EndPlace:      "PlaceB",
//		DepartureTime: "2023-01-01",
//	}
//	queryInfoResp, err := cli.QueryInfo(tripInfo)
//	if err != nil {
//		t.Errorf("QueryInfo request failed, err %s", err)
//	}
//	if len(queryInfoResp.Data) == 0 {
//		t.Errorf("QueryInfo returned no results")
//	}
//
//	// Test QueryInfoInParallel
//	queryInfoInParallelResp, err := cli.QueryInfoInParallel(tripInfo)
//	if err != nil {
//		t.Errorf("QueryInfoInParallel request failed, err %s", err)
//	}
//	if len(queryInfoInParallelResp.Data) == 0 {
//		t.Errorf("QueryInfoInParallel returned no results")
//	}
//
//	// Test GetTripAllDetailInfo
//	tripAllDetailInfo := TripAllDetailInfo{
//		TripId: getId,
//	}
//	tripAllDetailResp, err := cli.GetTripAllDetailInfo(tripAllDetailInfo)
//	if err != nil {
//		t.Errorf("GetTripAllDetailInfo request failed, err %s", err)
//	}
//	if tripAllDetailResp == nil {
//		t.Errorf("GetTripAllDetailInfo returned no result")
//	}
//
//	// Test AdminQueryAll
//	adminQueryAllResp, err := cli.AdminQueryAll()
//	if err != nil {
//		t.Errorf("AdminQueryAll request failed, err %s", err)
//	}
//	if len(adminQueryAllResp.Data) == 0 {
//		t.Errorf("AdminQueryAll returned no results")
//	}
//}
