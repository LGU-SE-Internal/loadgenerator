package service

import (
	"fmt"
	"testing"
)

func TestSvcImpl_ReqGetByCheapest(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqGetByCheapest(&TravelQueryInfo{
		DepartureTime: "2006-07-19",
		EndPlace:      "californiaairport",
		StartPlace:    "shenzhenbei",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqGetByMinStation(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqGetByMinStation(&TravelQueryInfo{
		DepartureTime: "2006-07-19",
		EndPlace:      "californiaairport",
		StartPlace:    "shenzhenbei",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqGetByQuickest(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqGetByQuickest(&TravelQueryInfo{
		DepartureTime: "2006-07-19",
		EndPlace:      "californiaairport",
		StartPlace:    "shenzhenbei",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqTransferResult(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqTransferResult(&TransferTravelQueryInfo{
		EndStation:   "北京",
		StartStation: "上海",
		TrainType:    "G",
		TravelDate:   "2023-04-01",
		ViaStation:   "南京",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}
