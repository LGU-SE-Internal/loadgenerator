package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"testing"
)

func TestCreateFoodDeliveryOrder(t *testing.T) {
	cli, _ := GetAdminClient()

	AddResp, err := cli.ReqCreateFoodDeliveryOrder(&FoodDeliveryOrder{
		CreatedTime:  faker.Date(),
		DeliveryFee:  20,
		DeliveryTime: faker.Date(),
		FoodList: []Food{{
			FoodName: "Hamburger",
			Price:    5.0,
		}},
		Id:                 uuid.NewString(),
		SeatNo:             RandomIntBetween(1, 6),
		StationFoodStoreId: "fc212d9b-4215-40ab-bc66-a02710fd387b",
		TripId:             uuid.NewString(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqFindAllFoodDeliveryOrders(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqGetAllFoodDeliveryOrders()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqGetFoodDeliveryOrderByStoreId(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqGetFoodDeliveryOrderByStoreId("fc212d9b-4215-40ab-bc66-a02710fd387b")
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqGetFoodDeliveryOrderById(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqGetFoodDeliveryOrderById("8a80811d9031564e0190366bc1950000")
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqDeleteFoodDeliveryOrderById(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqDeleteFoodDeliveryOrderById("8a80811d9031564e019036718aab0004")
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqUpdateDeliveryTime(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqUpdateDeliveryTime(&DeliveryInfo{
		DeliveryTime: faker.TimeString(),
		OrderId:      "8a80811d9031564e0190366bc1950000",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqUpdateSeatNo(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqUpdateSeatNo(&SeatInfo{
		SeatNo:  0,
		OrderId: "8a80811d9031564e0190366bc1950000",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqUpdateTripId(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqUpdateTripId(&TripOrderInfo{
		TripId:  uuid.NewString(),
		OrderId: "8a80811d9031564e0190366bc1950000",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}
