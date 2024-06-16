package service

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestStationService_FullIntegration(t *testing.T) {
	// Admin Test
	// Query Test
	cli, _ := GetAdminClient()
	resp, err := cli.QueryStations()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("resp.Status != 1")
	}

	//Mock
	MockedID := faker.UUIDHyphenated()
	input := &Station_station{
		ID:       MockedID,
		Name:     "Shenzhen Bei",
		StayTime: 7,
	}

	// Create Test
	resp1, err1 := cli.CreateStation(input)
	if err1 != nil {
		t.Errorf("Request failed, err1 %s", err1)
	}
	if resp1.Status != 1 {
		t.Errorf("Already exists")
	}

	// Test Update
	input1 := &Station_station{
		ID:       "c40200a8-bb63-445f-b332-6c4891666829",
		Name:     "zhenjiang",
		StayTime: 3,
	}
	resp2, err2 := cli.UpdateStation(input1)
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if resp2.Status != 1 {
		t.Errorf("resp2.Status != 1")
	}

	// Test Deletion
	var stationId string
	if len(resp.Data) > 0 {
		stationId = resp.Data[len(resp.Data)-1].Id
	} else {
		t.Errorf("resp.Data is empty")
	}
	stationId_delete := stationId
	//stationId := "45dea90e-eb9b-4602-8562-0b4dfdf12e5f"
	resp3, err3 := cli.DeleteStation(stationId_delete)
	if err3 != nil {
		t.Errorf("Request failed, err3 %s", err3)
	}
	if resp3.Status != 1 {
		t.Errorf("resp3.Status != 1")
	}

	// Test Query By name
	stationName := "beijing"
	resp4, err4 := cli.QueryStationIdByName(stationName)
	if err4 != nil {
		t.Errorf("Request failed, err4 %s", err4)
	}
	if resp4.Status != 1 {
		t.Errorf("resp4.Status != 1")
	}

	// Test Query by names
	stationNames := []string{"suzhou", "shijiazhuang"}
	resp5, err5 := cli.QueryStationIdsByNames(stationNames)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	if resp5.Status != 1 {
		t.Errorf("resp5.Status != 1")
	}

	// Test Query Name by ID
	var getStationId string
	if len(resp.Data) > 0 {
		getStationId = resp.Data[0].Id
	} else {
		t.Errorf("resp.Data is empty")
	}
	stationID := getStationId
	resp6, err6 := cli.QueryStationNameById(stationID)
	if err6 != nil {
		t.Errorf("Request failed, err6 %s", err6)
	}
	if resp6.Status != 1 {
		t.Errorf("resp6.Status != 1")
	}

	// Test QueryStationNamesByIds
	var id1 string
	var id2 string
	if len(resp.Data) > 1 {
		id1 = resp.Data[0].Id
		id2 = resp.Data[1].Id
	}
	stationIds := []string{id1, id2}
	resp7, err7 := cli.QueryStationNamesByIds(stationIds)
	if err7 != nil {
		t.Errorf("Request failed, err7 %s", err7)
	}
	if resp7.Status != 1 {
		t.Errorf("resp7.Status != 1")
	}
}
