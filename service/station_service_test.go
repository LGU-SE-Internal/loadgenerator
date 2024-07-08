package service

import (
	"github.com/go-faker/faker/v4"
	"math/rand"
	"strings"
	"testing"
)

func TestStationService_FullIntegration(t *testing.T) {
	// Admin Test
	// Query Test
	cli, _ := GetAdminClient()
	var stationSvc StationService = cli

	resp, err := stationSvc.QueryStations()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if len(resp.Data) == 0 {
		t.Errorf("No stations found")
	}
	t.Log(resp)

	//Mock
	MockedCityName := faker.GetRealAddress().City
	input := &Station{
		Name:     MockedCityName,
		StayTime: rand.Intn(30),
	}

	// Create Test
	resp1, err1 := stationSvc.CreateStation(input)
	if err1 != nil {
		t.Errorf("Request failed, err1 %s", err1)
	}
	if resp1.Msg == "Already exists" {
		t.Log("station found, skip")
		t.Skip()
	}
	if resp1.Data.Name != strings.Replace(strings.ToLower(input.Name), " ", "", -1) {
		t.Errorf("Request failed, resp1.Data.Name: %s, expected: %s", resp1.Data.Name, strings.Replace(strings.ToLower(input.Name), " ", "", -1))
	}
	if resp1.Data.StayTime != input.StayTime {
		t.Errorf("Request failed, resp1.Data.StayTime: %d, expected: %d", resp1.Data.StayTime, input.StayTime)
	}
	existedStation := resp1.Data

	// Query all
	QueryAll, err7 := stationSvc.QueryStations()
	if err7 != nil {
		t.Errorf("Request failed, err7 %s", err7)
	}
	found := false
	for _, station := range QueryAll.Data {
		if station.Name == existedStation.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("Request failed, station not found")
	}

	// Test Update
	input1 := &Station{}
	input1.StayTime = rand.Intn(30)
	input1.ID = existedStation.Id
	input1.Name = existedStation.Name
	resp2, err2 := stationSvc.UpdateStation(input1)
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if resp2.Status != 1 {
		t.Errorf("resp2.Status != 1")
	}
	if resp2.Data.StayTime != input1.StayTime {
		t.Errorf("Request failed. Expected %d, got %d", input1.StayTime, resp2.Data.StayTime)
	}

	// Test Query By name
	// Get name by Query
	resp4, err4 := stationSvc.QueryStationIdByName(existedStation.Name)
	if err4 != nil {
		t.Errorf("Request failed, err4 %s", err4)
	}
	if resp4.Status != 1 {
		t.Errorf("resp4.Status != 1")
	}
	if resp4.Data != existedStation.Id {
		t.Errorf("resp4.Data != input.ID, expected: '%s', actual: '%s'", existedStation.Id, resp4.Data)
	}
	t.Logf("Query By name response: %v", resp4)

	// Test Query by names
	stationNames := []string{existedStation.Name}
	resp5, err5 := stationSvc.QueryStationIdsByNames(stationNames)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	if resp5.Status != 1 {
		t.Errorf("resp5.Status != 1")
	}
	if len(resp5.Data) != len(stationNames) {
		t.Errorf("len(resp5.Data) != len(stationNames): %d, expected: %d", len(resp5.Data), len(stationNames))
	}

	// Test Query Name by ID
	resp6, err6 := stationSvc.QueryStationNameById(existedStation.Id)
	if err6 != nil {
		t.Errorf("Request failed, err6 %s", err6)
	}
	if resp6.Status != 1 {
		t.Errorf("resp6.Status != 1")
	}
	if resp6.Data != existedStation.Name {
		t.Errorf("resp6.Data != input.Name, expected: '%s', actual: '%s'", existedStation.Name, resp6.Data)
	}

	// Test QueryStationNamesByIds
	stationIds := []string{existedStation.Id}
	resp7, err7 := stationSvc.QueryStationNamesByIds(stationIds)
	if err7 != nil {
		t.Errorf("Request failed, err7 %s", err7)
	}
	if resp7.Status != 1 {
		t.Errorf("resp7.Status != 1")
	}
	found = false
	for _, name := range resp7.Data {
		if name == existedStation.Name {
			found = true
		}
	}
	if !found {
		t.Errorf("Request failed, station not found")
	}

	resp3, err3 := stationSvc.DeleteStation(existedStation.Id)
	if err3 != nil {
		t.Errorf("Request failed, err3 %s", err3)
	}
	if resp3.Status != 1 {
		t.Errorf("resp3.Status != 1")
	}

}
