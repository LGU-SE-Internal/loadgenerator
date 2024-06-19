package service

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestRouteService_FullIntegration(t *testing.T) {
	cli, _ := GetBasicClient()
	MockedID := faker.UUIDHyphenated()
	input := &RouteInfo{
		ID:           MockedID,
		StartStation: "Shenzhen Bei",
		EndStation:   "Jiulong Xi",
		StationList:  "Shenzhen Bei,Shkou,Jiulong Xi",
		DistanceList: "77,66,55",
	}
	resp, err := cli.CreateAndModifyRoute(input)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("resp.Status != 1")
	}

	//Test Query
	AllRoutes_By_Query, err2 := cli.QueryAllRoutes()
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if AllRoutes_By_Query.Status != 1 {
		t.Errorf("AllRoutes_By_Query.Status != 1")
	}

	//Test Query_By_Id
	routeId_Query := MockedID
	resp3, err3 := cli.QueryRouteById(routeId_Query)
	if err3 != nil {
		t.Errorf("Request failed, err3 %s", err3)
	}
	if resp3.Status != 1 {
		t.Errorf("resp3.Status != 1")
	}

	//Test Query_By_Ids
	var routeIds []string
	if len(AllRoutes_By_Query.Data) > 0 {
		routeIds = append(routeIds, AllRoutes_By_Query.Data[0].Id)
		routeIds = append(routeIds, AllRoutes_By_Query.Data[len(AllRoutes_By_Query.Data)-1].Id)
	} else {
		t.Errorf("AllRoutes_By_Query.Data is empty")
	}

	resp4, err4 := cli.QueryRoutesByIds(routeIds)
	if err4 != nil {
		t.Errorf("Request failed, err4 %s", err4)
	}
	if resp4.Status != 1 {
		t.Errorf("resp4.Status != 1")
	}

	// Test Deletion
	routeId := MockedID
	resp1, err1 := cli.DeleteRoute(routeId)
	if err1 != nil {
		t.Errorf("Request failed, err %s", err1)
	}
	if resp1.Status != 1 {
		t.Errorf("resp.Status != 1")
	}

	//Test Find by start and end
	start := "shanghai"
	end := "taiyuan"
	resp5, err5 := cli.QueryRoutesByStartAndEnd(start, end)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	t.Logf("QueryRoutesByStartAndEnd retuens: %v", resp5)

	// Test unshown data
	start_false := "California"
	end_false := "San Diego"
	resp6, err6 := cli.QueryRoutesByStartAndEnd(start_false, end_false)
	if err6 != nil {
		t.Errorf("Request failed, err6 %s", err6)
	}
	if resp6.Status != 0 {
		t.Errorf("resp6.Status != 0")
	}

}
