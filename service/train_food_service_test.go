package service

import "testing"

func TestSvcImplGetAllTrainFood(t *testing.T) {
	cli, _ := GetBasicClient()
	resp, err := cli.GetAllTrainFood()
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

	for _, train := range resp.Data {
		trainFood, err := cli.GetTrainFoodByTripId(train.TripId)
		if err != nil {
			t.Error(err)
		}
		t.Log(trainFood)
	}
}
