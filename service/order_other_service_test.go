package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestSvcImpl_ReqFindAllOeder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqFindAllOrderOther()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqCreateNewOeder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqCreateNewOrderOther(&Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
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

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqSaveOrderInfo_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqSaveOrderInfoOther(&Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "ee628cb0-6512-4dd0-ba1e-6eb5ccededaa",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 1,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqAddCreateNewOrder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqAddCreateNewOrderOther(&Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
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

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(AddResp.Msg)
}

func TestSvcImpl_ReqUpdateOrder_OrderService_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqUpdateOrderOrderServiceOther(&Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "ee628cb0-6512-4dd0-ba1e-6eb5ccededaa",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqPayOrder_Other(t *testing.T) {
	cli, _ := GetBasicClient()
	Resp, err := cli.ReqPayOrderOther("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetOrderPrice_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetOrderPriceOther("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqQueryOrders_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqQueryOrdersOther(&OrderInfo{
		BoughtDateEnd:         faker.Date(),
		BoughtDateStart:       faker.Date(),
		EnableBoughtDateQuery: true,
		EnableStateQuery:      true,
		EnableTravelDateQuery: true,
		LoginId:               uuid.NewString(),
		State:                 0,
		TravelDateEnd:         faker.Date(),
		TravelDateStart:       faker.Date(),
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqQueryOrderForRefresh_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqQueryOrderForRefreshOther(&OrderInfo{
		BoughtDateEnd:         faker.Date(),
		BoughtDateStart:       faker.Date(),
		EnableBoughtDateQuery: true,
		EnableStateQuery:      true,
		EnableTravelDateQuery: true,
		LoginId:               uuid.NewString(),
		State:                 0,
		TravelDateEnd:         faker.Date(),
		TravelDateStart:       faker.Date(),
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqSecurityInfoCheck_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqSecurityInfoCheckOther(faker.Date(), "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqModifyOrder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqModifyOrderOther("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa", 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetTicketsList_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetTicketsListOther(&Seat{
		DestStation:  RandomProvincialCapitalEN(),
		SeatType:     2,
		StartStation: RandomProvincialCapitalEN(),
		Stations:     nil,
		TotalNum:     0,
		TrainNumber:  GenerateTrainNumber(),
		TravelDate:   faker.Date(),
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqCalculateSoldTicket_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqCalculateSoldTicketOther(faker.Date(), GenerateTrainNumber())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetOrderById_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetOrderByIdOther("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqDeleteOrder_OrderService_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqDeleteOrderOrderServiceOther("4a766f5d-ef7c-4629-aef4-04492e247503")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_End2End_OrderOtherService(t *testing.T) {
	cli, _ := GetAdminClient()
	var orderSvc OrderOtherService = cli

	Resp, err := orderSvc.ReqFindAllOrderOther()
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
		SeatNumber:             rand.Intn(30),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp0, err := orderSvc.ReqCreateNewOrderOther(&originOrder0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder0 := Resp0.Data
	if !compareOrders(&originOrder0, &returnedOrder0) {
		t.Skip()
	}

	Resp1, err := orderSvc.ReqGetOrderByIdOther(returnedOrder0.Id)
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
		SeatNumber:             rand.Intn(30),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp2, err := orderSvc.ReqSaveOrderInfoOther(&originOrder1)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder2 := Resp2.Data

	if !compareOrders(&originOrder1, &returnedOrder2) {
		t.Skip()
	}

	Resp3, err := orderSvc.ReqGetOrderByIdOther(returnedOrder2.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder3 := Resp3.Data
	if !compareOrders(&returnedOrder2, &returnedOrder3) {
		t.Skip()
	}

	Resp4, err := orderSvc.ReqGetOrderPriceOther(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp4)

	Resp5, err := orderSvc.ReqPayOrderOther(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp5)

	Resp7, err := orderSvc.ReqDeleteOrderOrderServiceOther(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp7)
}

func TestSvcImpl_End2End_OrderOtherService_another(t *testing.T) {
	cli, _ := GetAdminClient()
	var orderSvc OrderOtherService = cli

	randomOrder := getRandomOrder_Other()
	prevBDay, nextBDay, err := getAdjacentDates(randomOrder.BoughtDate)
	prevTDay, nextTDay, err := getAdjacentDates(randomOrder.TravelDate)

	Resp6, err := orderSvc.ReqQueryOrdersOther(&OrderInfo{
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
	if len(Resp6.Data) < 1 {
		t.Errorf("[queryOrders] no orders found")
	}

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
		SeatNumber:             rand.Intn(30),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp8, err := orderSvc.ReqCreateNewOrderOther(&origin_order_0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	if !compareOrders(&origin_order_0, &Resp8.Data) {
		t.Errorf("【ReqCreateNewOrder】unexpected returned order.")
		t.Skip()
	}

	originOrder := Resp8.Data
	randomOrder = getRandomOrder_Other()
	prevBDay, nextBDay, err = getAdjacentDates(randomOrder.BoughtDate)
	prevTDay, nextTDay, err = getAdjacentDates(randomOrder.TravelDate)

	Resp9, _ := orderSvc.ReqQueryOrderForRefreshOther(&OrderInfo{
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
		t.Skip()
	}
	if !compareOrders(&randomOrder, &Resp9.Data[0]) {
		t.Errorf("【ReqQueryOrderForRefresh】unexpected returned order.")
	}
	Resp, err := orderSvc.ReqSecurityInfoCheckOther(originOrder.BoughtDate, originOrder.AccountId)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	fmt.Println(Resp.Msg)
	fmt.Println(Resp.Data.OrderNumInLastOneHour)
	fmt.Println(Resp.Data.OrderNumOfValidOrder)
	Resp21, err := orderSvc.ReqModifyOrderOther(originOrder.Id, 0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	if !compareOrders(&originOrder, &Resp21.Data) {
		t.Errorf("【ReqModifyOrder】unexpected returned order.")
		t.Skip()
	}

	var stations []string

	// 设置随机数种子，以确保每次运行程序时都能得到不同的随机数
	rand.Seed(time.Now().UnixNano())

	// 生成一个[0, 1)之间的浮点数
	randomFloat := rand.Float64()

	// 如果随机数小于0.5，则执行if代码块；否则，执行else代码块
	if randomFloat < 0.5 {
		stations = []string{randomOrder.To, randomOrder.From}
	} else {
		stations = []string{randomOrder.From, randomOrder.To}
	}

	Resp22, err := orderSvc.ReqGetTicketsListOther(&Seat{
		DestStation:  randomOrder.To,
		SeatType:     randomOrder.SeatClass,
		StartStation: randomOrder.From,
		Stations:     stations,
		TotalNum:     rand.Intn(10),
		TrainNumber:  randomOrder.TrainNumber,
		TravelDate:   randomOrder.TravelDate,
	})

	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	fmt.Println(Resp22.Msg)

	Resp23, _ := orderSvc.ReqCalculateSoldTicketOther(faker.Date(), GenerateTrainNumber())
	fmt.Println(Resp23.Msg)

	Resp24, _ := orderSvc.ReqGetOrderByIdOther(originOrder.Id)
	fmt.Println(Resp24.Msg)
	if !compareOrders(&originOrder, &Resp24.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}

	Resp25, _ := orderSvc.ReqDeleteOrderOrderServiceOther(originOrder.Id)
	fmt.Println(Resp25.Msg)
	if !compareOrders(&originOrder, &Resp25.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}

	randomContact = getRandomContact()
	newOrder := Order{
		AccountId:              randomContact.AccountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           randomContact.Name,
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     Resp8.Data.Id,
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             rand.Intn(30),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp26, _ := orderSvc.ReqAddCreateNewOrderOther(&newOrder)
	fmt.Println(Resp26.Msg)
	if !compareOrders(&newOrder, &Resp26.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}
	newOrder = Resp26.Data
	Resp27, _ := orderSvc.ReqUpdateOrderOrderServiceOther(&newOrder)
	fmt.Println(Resp27.Msg)
	if !compareOrders(&newOrder, &Resp27.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}
	Resp28, _ := orderSvc.ReqDeleteOrderOrderServiceOther(Resp27.Data.Id)
	fmt.Println(Resp28.Msg)
}
