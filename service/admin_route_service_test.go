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

	loginID := faker.UUIDHyphenated()
	id := faker.UUIDHyphenated()

	// Test AddRoute for Admin User
	addRouteResp, err := adminClient.ReqAddRoute(&RouteInfo{
		LoginID:      loginID,
		StartStation: "shenzhenbei",
		EndStation:   "jiulong",
		StationList:  "shenzhenbei,shekou,jiulong",
		DistanceList: "100, 50, 10",
		ID:           id,
	})
	if err != nil {
		t.Errorf("Failed to add route for admin user: %v", err)
	}
	if addRouteResp.Status != 200 {
		t.Errorf("Expected status 200, got %d", addRouteResp.Status)
	}

	// Extract route ID from the response
	var routeID string
	if len(allRoutesResp.Data) > 0 {
		routeID = allRoutesResp.Data[len(allRoutesResp.Data)-1].ID
	} else {
		t.Errorf("allRoutesResp.Data is empty")
	}

	// Test DeleteRoute for Admin User
	deleteRouteResp, err := adminClient.ReqDeleteRoute(routeID)
	if err != nil {
		t.Errorf("Failed to delete route for admin user: %v", err)
	}
	if deleteRouteResp.Status != 200 {
		t.Errorf("Expected status 200, got %d", deleteRouteResp.Status)
	}

	// Create a client for basic user
	basicClient, _ := GetBasicClient()

	// Test GetAllRoutes for Basic User
	allRoutesResp, err = basicClient.ReqGetAllRoutes()
	if err != nil {
		t.Errorf("Failed to get all routes: %v", err)
	}

	loginID_basic := faker.UUIDHyphenated()
	id_basic := faker.UUIDHyphenated()

	// Test AddRoute for Basic User
	addRouteResp, err = basicClient.ReqAddRoute(&RouteInfo{
		LoginID:      loginID_basic,
		StartStation: "Jiu Long",
		EndStation:   "Shenzhen",
		StationList:  "Jiu Long, She Kou, Shenzhen",
		DistanceList: "10, 50, 100",
		ID:           id_basic,
	})
	if err != nil {
		t.Errorf("Failed to add route for basic user: %v", err)
	}
	if addRouteResp.Status != 200 {
		t.Errorf("Expected status 200, got %d", addRouteResp.Status)
	}

	// Extract route ID from the response
	if len(allRoutesResp.Data) > 0 {
		routeID = allRoutesResp.Data[len(allRoutesResp.Data)-1].ID
	} else {
		t.Errorf("allRoutesResp.Data is empty")
	}

	// Test DeleteRoute for Basic User
	deleteRouteResp, err = basicClient.ReqDeleteRoute(routeID)
	if err != nil {
		t.Errorf("Failed to delete route for basic user: %v", err)
	}
	if deleteRouteResp.Status != 200 {
		t.Errorf("Expected status 200, got %d", deleteRouteResp.Status)
	}
}
