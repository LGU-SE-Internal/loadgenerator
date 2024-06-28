package service

import (
	"testing"
)

func TestPlanService_FullIntegration(t *testing.T) {
	//Admin
	cli, _ := GetAdminClient()

	//Mock
	var routeSvc RouteService = cli
	var MockedStartStation string
	var MockedEndStation string
	QueryAllRoute, err := routeSvc.QueryAllRoutes()
	if err != nil {
		t.Error(err.Error())
	}

	if len(QueryAllRoute.Data) > 0 {
		MockedStartStation = QueryAllRoute.Data[0].StartStation
		MockedEndStation = QueryAllRoute.Data[0].EndStation
	}
	input := &RoutePlanInfo{
		StartStation: MockedStartStation,
		EndStation:   MockedEndStation,
		Num:          3,
		TravelDate:   "2024-07-04 12:00:00",
	}

	// Test GetCheapestRoutes
	resp, err := cli.GetCheapestRoutes(input)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("resp.Status != 1")
	}

	// Test GetQuickestRoutes
	resp2, err2 := cli.GetQuickestRoutes(input)
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if resp2.Status != 1 {
		t.Errorf("resp2.Status != 1")
	}

	// Test GetMinStopStations
	resp3, err3 := cli.GetMinStopStations(input)
	if err3 != nil {
		t.Errorf("Request failed, err3 %s", err3)
	}
	t.Logf("GetMinStopStations returns: %v", resp3)
}
