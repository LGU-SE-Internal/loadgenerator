package service

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestAdminRouteService_FullIntegration(t *testing.T) {
	// Create a client for admin user
	adminClient, _ := GetAdminClient()

	// Test GetAllRoutes for Admin User
	allRoutesResp, err := adminClient.ReqGetAllRoutes()
	if err != nil {
		t.Errorf("Failed to get all routes: %v", err)
	}
	if allRoutesResp.Status != 1 {
		t.Errorf("resp.Status != 1")
	}

	// Mock
	loginID := faker.UUIDHyphenated()
	id := faker.UUIDHyphenated()
	input := &RouteInfo{
		LoginID:      loginID,
		StartStation: "Shenzhen Bei",
		EndStation:   "Jiulong Xi",
		StationList:  "Shenzhen Bei,Jiulong Xi",
		DistanceList: "100, 50, 10",
		ID:           id,
	}

	// Test AddRoute for Admin User
	addRouteResp, err := adminClient.ReqAddRoute(input)
	if err != nil {
		t.Errorf("Failed to add route for admin user: %v", err)
	}
	t.Logf("AddRoute for Admin User:  %v", addRouteResp.Msg)

	// Extract route ID from the response
	var routeID string
	if len(allRoutesResp.Data) > 0 {
		routeID = allRoutesResp.Data[len(allRoutesResp.Data)-1].Id
	} else {
		t.Errorf("allRoutesResp.Data is empty")
	}

	// Test DeleteRoute for Admin User
	deleteRouteResp, err := adminClient.ReqDeleteRoute(routeID)
	if err != nil {
		t.Errorf("Failed to delete route for admin user: %v", err)
	}
	if deleteRouteResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", deleteRouteResp.Status)
	}
}
