package service

import (
	"fmt"
	"testing"
)

func TestSvcImpl_ReqGetByCheapest(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqGetByCheapest(&TravelQueryInfo{
		DepartureTime: "2024-07-19",
		EndPlace:      "taiyuan",
		StartPlace:    "shanghai",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqGetByMinStation(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqGetByMinStation(&TravelQueryInfo{
		DepartureTime: "2024-07-19",
		EndPlace:      "taiyuan",
		StartPlace:    "shanghai",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqGetByQuickest(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqGetByQuickest(&TravelQueryInfo{
		DepartureTime: "2024-07-19",
		EndPlace:      "taiyuan",
		StartPlace:    "shanghai",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqTransferResult(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqTransferResult(&TransferTravelQueryInfo{
		EndStation:   "taiyuan",
		StartStation: "nanjing",
		TrainType:    "1",
		TravelDate:   "2024-07-01",
		ViaStation:   "shanghai",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}
