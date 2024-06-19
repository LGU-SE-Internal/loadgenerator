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
		DeliveryFee:  RandomIntBetween(1, 10),
		DeliveryTime: faker.Date(),
		FoodList: []Food{{
			FoodName: getRandomDish(),
			Price:    RandomIntBetween(1, 10),
		}},
		Id:                 uuid.NewString(),
		SeatNo:             RandomIntBetween(1, 6),
		StationFoodStoreId: uuid.NewString(),
		TripId:             uuid.NewString(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}
