package service

import "testing"

func TestSvcImplGetAllTrainFood(t *testing.T) {
	cli, _ := GetBasicClient()
	var trainFoodSvc TrainFoodService = cli

	resp, err := trainFoodSvc.GetAllTrainFood()
	if err != nil {
		t.Error(err)
	}
	if resp.Status != 1 {
		t.Errorf("GetAllTrainFood's status should be 1 but got %d", resp.Status)
	}
	t.Log(resp)

	for _, train := range resp.Data {
		trainFood, err := trainFoodSvc.GetTrainFoodByTripId(train.TripId)
		if err != nil {
			t.Error(err)
		}
		if trainFood.Status != 1 {
			t.Errorf("GetTrainFoodByTripId's status should be 1 but got %d", trainFood.Status)
		}
		t.Log(trainFood)
	}
}
