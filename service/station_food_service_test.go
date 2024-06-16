package service

import "testing"

func TestSvcImpl_GetAllStationFood(t *testing.T) {
	cli, _ := GetAdminClient()

	resp, err := cli.GetAllStationFood()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

	for _, stationFood := range resp.Data {
		sf, err := cli.GetStationFoodByNames([]string{stationFood.StationName})
		if err != nil {
			t.Error(err)
		}
		t.Log(sf)

		sf1, err := cli.GetStationFoodByName(stationFood.StationName)
		if err != nil {
			t.Error(err)
		}
		t.Log(sf1)

		sf2, err := cli.GetStationFoodById(stationFood.Id)
		if err != nil {
			t.Error(err)
		}

		t.Log(sf2)
	}

}
