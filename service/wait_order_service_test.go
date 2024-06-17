package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"testing"
)

func TestSvcImpl_ReqCreateNewOrder_WaitOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqCreateNewWaitOrder(&OrderVO{
		AccountId:  uuid.NewString(),
		ContactsId: uuid.NewString(),
		Date:       faker.Date(),
		From:       RandomProvincialCapitalEN(),
		Price:      RandomDecimalStringBetween(1, 10),
		SeatType:   RandomIntBetween(0, 1),
		To:         RandomProvincialCapitalEN(),
		TripId:     uuid.NewString(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqGetAllOrders_WaitOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqGetAllWaitOrder()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqGetWaitListOrders(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqGetWaitListOrders()
	fmt.Println(GetResp.Msg)
}
