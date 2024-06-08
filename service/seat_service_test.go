package service

import "testing"

func TestSvcImpl_ReqSeatCreate(t *testing.T) {
	cli, _ := GetBasicClient()
	resp, err := cli.ReqSeatCreate(&SeatCreateInfoReq{
		TravelDate:  "2024-06-06 14:16:00",
		TrainNumber: "777",
		DestStation: "shenzhen",
		SeatType:    7,
		TotalNum:    8,
		Stations:    []string{"shenzhen", "suzhou", "hong kong"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", resp)
}
