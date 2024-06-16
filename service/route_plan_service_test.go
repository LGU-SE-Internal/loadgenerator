package service

import (
	"testing"
)

func TestPlanService_FullIntegration(t *testing.T) {
	//Admin
	cli, _ := GetAdminClient()

	//Mock
	input := &RoutePlanInfo{
		StartStation: "Shenzhen Bei",
		EndStation:   "Jiulong Xi",
		Num:          1,
		TravelDate:   "2024-06-14",
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
	if resp3.Status != 1 {
		t.Errorf("resp3.Status != 1")
	}
}
