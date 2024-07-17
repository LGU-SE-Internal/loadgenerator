package service

import "testing"

func TestSvcImpl_GetAllStationFood(t *testing.T) {
	cli, _ := GetAdminClient()
	var stationFoodSvc StationFoodService = cli

	resp, err := stationFoodSvc.GetAllStationFood()
	if err != nil {
		t.Error(err)
	}
	if resp.Status != 1 {
		t.Errorf("GetAllStationFood status should be 1, but is %d", resp.Status)
	}
	t.Log(resp)

	for _, stationFood := range resp.Data {
		sf, err := stationFoodSvc.GetStationFoodByNames([]string{stationFood.StationName})
		if err != nil {
			t.Error(err)
		}
		if sf.Status != 1 {
			t.Errorf("GetStationFoodByNames status should be 1, but is %d", sf.Status)
		}
		t.Log(sf)

		sf1, err := stationFoodSvc.GetStationFoodByName(stationFood.StationName)
		if err != nil {
			t.Error(err)
		}
		if sf1.Status != 1 {
			t.Errorf("GetStationFoodByName status should be 1, but is %d", sf1.Status)
		}
		t.Log(sf1)

		sf2, err := stationFoodSvc.GetStationFoodById(stationFood.Id)
		if err != nil {
			t.Error(err)
		}
		if sf2.Status != 1 {
			t.Errorf("GetStationFoodById status should be 1, but is %d", sf2.Status)
		}
		t.Log(sf2)
	}

}
