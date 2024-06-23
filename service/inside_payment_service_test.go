package service

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestSvcImpl_ReqPay_InsidePayment(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqPay_InsidePayment(&TripPayment{
		TripId:  "G111",
		OrderId: "f1d1660a-bfb8-4304-9abe-018fef31a484",
		Price:   "19",
		UserId:  "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqCreateAccount(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqCreateAccount(&AccountInfo{
		Money:  RandomDecimalStringBetween(1, 100),
		UserId: uuid.NewString(),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqPayDifference(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqPayDifference(&TripPayment{
		TripId:  "G111",
		OrderId: "f1d1660a-bfb8-4304-9abe-018fef31a484",
		Price:   "19",
		UserId:  "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqPay_InsidePayment2(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqQueryAccount()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqQueryAccount(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqQueryAccount()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqDrawBack(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqDrawBack("4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f", "12")
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqQueryAddMoney(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqQueryAddMoney()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqQueryInsidePayment(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqQueryInsidePayment()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqAddMoney_Inside(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqAddMoney_Inside("4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f", "12")
	fmt.Println(GetResp.Msg)
}
