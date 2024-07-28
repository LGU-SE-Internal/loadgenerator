package service

import (
	//"github.com/Lincyaw/loadgenerator/behaviors"

	//"github.com/Lincyaw/loadgenerator/behaviors"
	"log"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestTravel2Service_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()

	var travelSvc TravelService = cli
	var travel2Svc Travel2Service = cli
	var MockedTripId string
	GetAllTravelInfo, err := travelSvc.QueryAllTrip()
	if err != nil {
		t.Errorf("error occurs: %v", err)
	}

	if len(GetAllTravelInfo.Data) > 0 {
		MockedTripId = GetAllTravelInfo.Data[0].Id
	} else {
		log.Fatalf("no travel info")
	}
	//MockedTripId := faker.UUIDHyphenated()
	//MockedRouteId := faker.UUIDHyphenated()
	//MockedTrainType := "HighSpeed"

	// Test GetTrainTypeByTripId
	trainTypeResp, err := travel2Svc.GetTrain2TypeByTripId(MockedTripId)
	if err != nil {
		t.Errorf("GetTrainTypeByTripId failed: %v", err)
	}
	t.Logf("GetTrainTypeByTripId response: %+v", trainTypeResp)

	// Test GetRouteByTripId
	routeResp, err := travel2Svc.GetRouteByTrip2Id(MockedTripId)
	if err != nil {
		t.Errorf("GetRouteByTripId failed: %v", err)
	}
	t.Logf("GetRouteByTripId response: %+v", routeResp)

	// Mock data
	//MockedTypeName := faker.Word()
	MockedTripID := GenerateTripId()
	MockedLoginId := faker.UUIDHyphenated()
	//MockedIndex := 1
	//MockedTripIDName := faker.Word()
	MockedTrainTypeName := faker.Word()
	MockedRouteID := faker.UUIDHyphenated()
	MockedStartStationName := "suzhou"
	MockedTerminalStationName := "taiyuan"
	MockedStartTime := faker.Date()
	MockedEndTime := faker.Date()

	travelInfo := TravelInfo{
		LoginID:             MockedLoginId,
		TripID:              MockedTripID,
		TrainTypeName:       MockedTrainTypeName,
		RouteID:             MockedRouteID,
		StartStationName:    MockedStartStationName,
		StationsName:        "suzhou, taiyuan",
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}
	// Test Create
	createResp, err := travel2Svc.CreateTrip2(&travelInfo)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}
	t.Logf("Create response: %+v", createResp)

	// Test Retrieve
	retrieveResp, err := travel2Svc.RetrieveTrip2(MockedTripId)
	if err != nil {
		t.Errorf("Retrieve failed: %v", err)
	}
	t.Logf("Retrieve response: %+v", retrieveResp)

	// Test Update
	updateTravelInfo := TravelInfo{
		LoginID:             MockedLoginId,
		TripID:              "G777",
		TrainTypeName:       MockedTrainTypeName,
		RouteID:             MockedRouteID,
		StartStationName:    MockedStartStationName,
		StationsName:        "suzhou, taiyuan",
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}
	updateResp, err := travel2Svc.UpdateTrip2(&updateTravelInfo)
	if err != nil {
		t.Errorf("Update failed: %v", err)
	}
	t.Logf("Update response: %+v", updateResp)

	// Test Delete
	deleteResp, err := travel2Svc.DeleteTrip2(MockedTripId)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	t.Logf("Delete response: %+v", deleteResp)

	// Test QueryByBatch
	queryByBatchReq := &TripInfo{
		StartPlace:    RandomProvincialCapitalEN(),
		EndPlace:      RandomProvincialCapitalEN(),
		DepartureTime: "2024-07-28 09:09:04",
	}
	queryByBatchResp, err := travel2Svc.QueryByBatch(queryByBatchReq)
	if err != nil {
		t.Errorf("QueryByBatch failed: %v", err)
	}
	t.Logf("QueryByBatch response: %+v", queryByBatchResp)

	// Test GetTripAllDetailInfo
	tripAllDetailInfo := Trip2AllDetailInfo{
		TripID:     "G777",
		TravelDate: "2024-06-28 09:09:04",
		From:       "suzhou",
		To:         "taiyuan",
	}
	tripDetailResp, err := travel2Svc.GetTrip2AllDetailInfo(&tripAllDetailInfo)
	if err != nil {
		t.Errorf("GetTripAllDetailInfo failed: %v", err)
	}
	t.Logf("GetTripAllDetailInfo response: %+v", tripDetailResp)

	// Test QueryAllTrip
	queryAllResp, err := travel2Svc.QueryAllTravel()
	if err != nil {
		t.Errorf("QueryAllTrip failed: %v", err)
	}
	t.Logf("QueryAllTrip response: %+v", queryAllResp)

	// Test AdminQueryAll
	adminQueryAllResp, err := travel2Svc.AdminQueryAllTravel()
	if err != nil {
		t.Errorf("AdminQueryAll failed: %v", err)
	}
	t.Logf("AdminQueryAll response: %+v", adminQueryAllResp)
}

// helper function
//func GenerateTripId() string {
//	// 设置随机数种子
//	rand.Seed(time.Now().UnixNano())
//
//	// 定义可能的开头字母
//	letters := []rune{'Z', 'T', 'K', 'G', 'D'}
//
//	// 随机选择一个字母
//	startLetter := letters[rand.Intn(len(letters))]
//
//	// 生成三个随机数字
//	randomNumber := rand.Intn(1000)
//
//	// 格式化成三位数字，不足三位前面补零
//	MockedTripID := fmt.Sprintf("%c%03d", startLetter, randomNumber)
//
//	return MockedTripID
//}
