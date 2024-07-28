package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"testing"
)

func TestRouteService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	var routeSvc RouteService = cli

	// Create
	MockedID := faker.UUIDHyphenated()
	MockedStartStation := faker.GetRealAddress().City
	MockedEndStation := faker.GetRealAddress().City
	MockedStationList := fmt.Sprintf("%s,%s,%s", MockedStartStation, faker.GetRealAddress().City, MockedEndStation)
	MockedDistanceList := fmt.Sprintf("%d,%d,%d", rand.Intn(30), rand.Intn(30), rand.Intn(30))
	input := RouteInfo{
		ID:           MockedID,
		StartStation: MockedStartStation,
		EndStation:   MockedEndStation,
		StationList:  MockedStationList,
		DistanceList: MockedDistanceList,
	}
	resp, err := routeSvc.CreateAndModifyRoute(&input)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Msg == "Already exists" {
		t.Log("Route already exists, skip")
		t.Skip()
	}
	if resp.Data.Id != input.ID {
		t.Errorf("Route ID does not match, expect %s, got %s", input.ID, resp.Data.Id)
	}
	if resp.Data.StartStation != input.StartStation {
		t.Errorf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
	}
	if resp.Data.EndStation != input.EndStation {
		t.Errorf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
	}
	if StringSliceToString(resp.Data.Stations) != ConvertCommaSeparatedToBracketed(input.StationList) {
		t.Errorf("StationList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.StationList), StringSliceToString(resp.Data.Stations))
	}
	if IntSliceToString(resp.Data.Distances) != ConvertCommaSeparatedToBracketed(input.DistanceList) {
		t.Errorf("DistanceList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.DistanceList), IntSliceToString(resp.Data.Distances))
	}
	existedRoute := resp.Data

	//Test Query
	AllRoutesByQuery, err2 := routeSvc.QueryAllRoutes()
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if AllRoutesByQuery.Status != 1 {
		t.Errorf("AllRoutes_By_Query.Status != 1")
	}
	found := false
	for _, route := range AllRoutesByQuery.Data {
		if route.Id == existedRoute.Id {
			found = true
		}
	}
	if !found {
		t.Errorf("Route not found by queryALL")
	}

	//Test Query_By_Id
	routeId_Query := existedRoute.Id
	resp3, err3 := routeSvc.QueryRouteById(routeId_Query)
	if err3 != nil {
		t.Errorf("Request failed, err3 %s", err3)
	}
	if resp3.Status != 1 {
		t.Errorf("resp3.Status != 1")
	}
	if resp3.Data.Id != existedRoute.Id {
		t.Errorf("resp3.Data.Id != existedRoute.Id, expect %s, got %s", existedRoute.Id, resp3.Data.Id)
	}

	//Test Query_By_Ids
	//var routeIds []string
	//if len(AllRoutes_By_Query.Data) > 0 {
	//	routeIds = append(routeIds, AllRoutes_By_Query.Data[0].Id)
	//	routeIds = append(routeIds, AllRoutes_By_Query.Data[len(AllRoutes_By_Query.Data)-1].Id)
	//} else {
	//	t.Errorf("AllRoutes_By_Query.Data is empty")
	//}

	routeIds := []string{existedRoute.Id, AllRoutesByQuery.Data[0].Id}
	resp4, err4 := routeSvc.QueryRoutesByIds(routeIds)
	if err4 != nil {
		t.Errorf("Request failed, err4 %s", err4)
	}
	if resp4.Status != 1 {
		t.Errorf("resp4.Status != 1")
	}
	found = false
	for _, route := range resp4.Data {
		if route.Id == existedRoute.Id {
			found = true
		}
	}
	if !found {
		t.Errorf("Route not found by query by QueryByStations")
	}

	//Test Find by start and end
	start := existedRoute.StartStation
	end := existedRoute.EndStation
	resp5, err5 := routeSvc.QueryRoutesByStartAndEnd(start, end)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	if resp5.Status != 1 {
		t.Errorf("resp5.Status != 1")
	}

	found = false
	for _, route := range resp5.Data {
		if route.Id == existedRoute.Id {
			found = true
		}
	}
	if !found {
		t.Errorf("Route not found by query by start and end")
	}
	t.Logf("QueryRoutesByStartAndEnd retuens: %v", resp5)

	// Test unshown data
	start_false := "California"
	end_false := "San Diego"
	resp6, err6 := routeSvc.QueryRoutesByStartAndEnd(start_false, end_false)
	if err6 != nil {
		t.Errorf("Request failed, err6 %s", err6)
	}
	if resp6.Status != 0 {
		t.Errorf("resp6.Status != 0")
	}

	// Test Deletion
	routeId := existedRoute.Id
	resp1, err1 := routeSvc.DeleteRoute(routeId)
	if err1 != nil {
		t.Errorf("Request failed, err %s", err1)
	}
	if resp1.Status != 1 {
		t.Errorf("resp.Status != 1")
	}

}
