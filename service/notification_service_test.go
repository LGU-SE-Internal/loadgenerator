package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"testing"
)

func TestSvcImpl_ReqOrderCancelSuccess(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqOrderCancelSuccess(&TicketOrder{
		Date:        randomTime(),
		Email:       faker.Email(),
		EndPlace:    RandomProvincialCapitalEN(),
		ID:          uuid.NewString(),
		OrderNumber: RandomDecimalStringBetween(1, 6),
		Price:       RandomDecimalStringBetween(2, 8),
		SeatClass:   RandomDecimalStringBetween(1, 3),
		SeatNumber:  GenerateSeatNumber(),
		SendStatus:  true,
		StartPlace:  RandomProvincialCapitalEN(),
		StartTime:   randomTime(),
		Username:    faker.Username(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*AddResp)
}

func TestSvcImpl_ReqOrderChangedSuccess(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqOrderChangedSuccess(&TicketOrder{
		Date:        randomTime(),
		Email:       faker.Email(),
		EndPlace:    RandomProvincialCapitalEN(),
		ID:          uuid.NewString(),
		OrderNumber: RandomDecimalStringBetween(1, 6),
		Price:       RandomDecimalStringBetween(2, 8),
		SeatClass:   RandomDecimalStringBetween(1, 3),
		SeatNumber:  GenerateSeatNumber(),
		SendStatus:  true,
		StartPlace:  RandomProvincialCapitalEN(),
		StartTime:   randomTime(),
		Username:    faker.Username(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*AddResp)
}

func TestSvcImpl_ReqOrderCreateSuccess(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqOrderCreateSuccess(&TicketOrder{
		Date:        randomTime(),
		Email:       faker.Email(),
		EndPlace:    RandomProvincialCapitalEN(),
		ID:          uuid.NewString(),
		OrderNumber: RandomDecimalStringBetween(1, 6),
		Price:       RandomDecimalStringBetween(2, 8),
		SeatClass:   RandomDecimalStringBetween(1, 3),
		SeatNumber:  GenerateSeatNumber(),
		SendStatus:  true,
		StartPlace:  RandomProvincialCapitalEN(),
		StartTime:   randomTime(),
		Username:    faker.Username(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*AddResp)
}

func TestSvcImpl_ReqPreserveSuccess(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqPreserveSuccess(&TicketOrder{
		Date:        randomTime(),
		Email:       faker.Email(),
		EndPlace:    RandomProvincialCapitalEN(),
		ID:          uuid.NewString(),
		OrderNumber: RandomDecimalStringBetween(1, 6),
		Price:       RandomDecimalStringBetween(2, 8),
		SeatClass:   RandomDecimalStringBetween(1, 3),
		SeatNumber:  GenerateSeatNumber(),
		SendStatus:  true,
		StartPlace:  RandomProvincialCapitalEN(),
		StartTime:   randomTime(),
		Username:    faker.Username(),
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*AddResp)
}

func TestSvcImpl_ReqTestSendMail(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqTestSendMail()
	fmt.Println(*GetResp)
}

func TestSvcImpl_ReqTestSend(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqTestSend()
	fmt.Println(*GetResp)
}
