package service

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

func TestSvcImpl_End2End_OrderService(t *testing.T) {
	cli, _ := GetAdminClient()
	var orderSvc OrderService = cli

	Resp, err := orderSvc.ReqFindAllOrder()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	if len(Resp.Data) == 0 {
		t.Log("no data found.")
	}

	randomContact := getRandomContact()
	originOrder0 := Order{
		AccountId:              randomContact.AccountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           randomContact.Name,
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "nil",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp0, err := orderSvc.ReqCreateNewOrder(&originOrder0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder0 := Resp0.Data
	if !compareOrders(&originOrder0, &returnedOrder0) {
		t.Skip()
	}

	Resp1, err := orderSvc.ReqGetOrderById(returnedOrder0.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder1 := Resp1.Data
	if !compareOrders(&returnedOrder0, &returnedOrder1) {
		t.Skip()
	}

	randomContact = getRandomContact()
	originOrder1 := Order{
		AccountId:              randomContact.AccountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           randomContact.Name,
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     returnedOrder0.Id,
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp2, err := orderSvc.ReqSaveOrderInfo(&originOrder1)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder2 := Resp2.Data

	if !compareOrders(&originOrder1, &returnedOrder2) {
		t.Skip()
	}

	Resp3, err := orderSvc.ReqGetOrderById(returnedOrder2.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder3 := Resp3.Data
	if !compareOrders(&returnedOrder2, &returnedOrder3) {
		t.Skip()
	}

	Resp4, err := orderSvc.ReqGetOrderPrice(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp4)

	Resp5, err := orderSvc.ReqPayOrder(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp5)

	Resp7, err := orderSvc.ReqDeleteOrder_OrderService(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp7)
}

func TestSvcImpl_End2End_OrderService_another(t *testing.T) {
	cli, _ := GetAdminClient()
	var orderSvc OrderService = cli

	randomOrder := getRandomOrder()
	prevBDay, nextBDay, err := getAdjacentDates(randomOrder.BoughtDate)
	prevTDay, nextTDay, err := getAdjacentDates(randomOrder.TravelDate)

	Resp6, err := orderSvc.ReqQueryOrders(&Qi{
		BoughtDateEnd:         nextBDay,
		BoughtDateStart:       prevBDay,
		EnableBoughtDateQuery: true,
		EnableStateQuery:      true,
		EnableTravelDateQuery: true,
		LoginId:               randomOrder.AccountId,
		State:                 0,
		TravelDateEnd:         nextTDay,
		TravelDateStart:       prevTDay,
	})

	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	t.Log(Resp6)

	randomContact := getRandomContact()
	origin_order_0 := Order{
		AccountId:              randomContact.AccountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           randomContact.Name,
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "nil",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp8, err := orderSvc.ReqCreateNewOrder(&origin_order_0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}

	originOrder := Resp8.Data
	randomOrder = getRandomOrder()
	prevBDay, nextBDay, err = getAdjacentDates(randomOrder.BoughtDate)
	prevTDay, nextTDay, err = getAdjacentDates(randomOrder.TravelDate)

	Resp9, _ := orderSvc.ReqQueryOrderForRefresh(&Qi{
		BoughtDateEnd:         nextBDay,
		BoughtDateStart:       prevBDay,
		EnableBoughtDateQuery: true,
		EnableStateQuery:      true,
		EnableTravelDateQuery: true,
		LoginId:               randomOrder.AccountId,
		State:                 0,
		TravelDateEnd:         nextTDay,
		TravelDateStart:       prevTDay,
	})
	fmt.Println(Resp9.Msg)
	Resp, _ := orderSvc.ReqSecurityInfoCheck(originOrder.BoughtDate, originOrder.AccountId)
	fmt.Println(Resp.Msg)
	Resp, _ = orderSvc.ReqModifyOrder(originOrder.Id, 0)
	fmt.Println(Resp.Msg)
	Resp, _ = orderSvc.ReqGetTicketsList(&Seat{
		DestStation:  RandomProvincialCapitalEN(),
		SeatType:     2,
		StartStation: RandomProvincialCapitalEN(),
		Stations:     nil,
		TotalNum:     0,
		TrainNumber:  GenerateTrainNumber(),
		TravelDate:   faker.Date(),
	})
	fmt.Println(Resp.Msg)
	Resp, _ = orderSvc.ReqCalculateSoldTicket(faker.Date(), GenerateTrainNumber())
	fmt.Println(Resp.Msg)
	Resp, _ = orderSvc.ReqGetOrderById(originOrder.Id)
	fmt.Println(Resp.Msg)
	Resp, _ = orderSvc.ReqDeleteOrder_OrderService(originOrder.Id)
	fmt.Println(Resp.Msg)
	randomContact = getRandomContact()
	Resp21, _ := orderSvc.ReqAddCreateNewOrder(&Order{
		AccountId:              randomContact.AccountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           randomContact.Name,
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     uuid.NewString(),
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	})
	fmt.Println(Resp21.Msg)
	Resp22, _ := orderSvc.ReqUpdateOrder_OrderService(&Resp21.Data)
	fmt.Println(Resp22.Msg)
	Resp23, _ := orderSvc.ReqDeleteOrder_OrderService(Resp22.Data.Id)
	fmt.Println(Resp23.Msg)
}
