package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
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
		SeatNumber:             rand.Intn(30),
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
		SeatNumber:             rand.Intn(30),
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

	Resp6, err := orderSvc.ReqQueryOrders(&OrderInfo{
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

	Resp8, err := orderSvc.ReqCreateNewOrder(&origin_order_0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	if !compareOrders(&origin_order_0, &Resp8.Data) {
		t.Errorf("【ReqCreateNewOrder】unexpected returned order.")
		t.Skip()
	}

	originOrder := Resp8.Data
	randomOrder = getRandomOrder()
	prevBDay, nextBDay, err = getAdjacentDates(randomOrder.BoughtDate)
	prevTDay, nextTDay, err = getAdjacentDates(randomOrder.TravelDate)

	orderInfo := OrderInfo{
		BoughtDateEnd:         nextBDay,
		BoughtDateStart:       prevBDay,
		EnableBoughtDateQuery: true,
		EnableStateQuery:      true,
		EnableTravelDateQuery: true,
		LoginId:               randomOrder.AccountId,
		State:                 0,
		TravelDateEnd:         nextTDay,
		TravelDateStart:       prevTDay,
	}

	Resp9, _ := orderSvc.ReqQueryOrderForRefresh(&orderInfo)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	if !compareOrders(&randomOrder, &Resp9.Data[0]) {
		t.Errorf("【ReqQueryOrderForRefresh】unexpected returned order.")
	}
	Resp, err := orderSvc.ReqSecurityInfoCheck(originOrder.BoughtDate, originOrder.AccountId)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	fmt.Println(Resp.Msg)
	fmt.Println(Resp.Data.OrderNumInLastOneHour)
	fmt.Println(Resp.Data.OrderNumOfValidOrder)
	Resp21, err := orderSvc.ReqModifyOrder(originOrder.Id, 0)
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

	Resp22, err := orderSvc.ReqGetTicketsList(&Seat{
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

	Resp23, _ := orderSvc.ReqCalculateSoldTicket(faker.Date(), GenerateTrainNumber())
	fmt.Println(Resp23.Msg)

	Resp24, _ := orderSvc.ReqGetOrderById(originOrder.Id)
	fmt.Println(Resp24.Msg)
	if !compareOrders(&originOrder, &Resp24.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}

	Resp25, _ := orderSvc.ReqDeleteOrder_OrderService(originOrder.Id)
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

	Resp26, _ := orderSvc.ReqAddCreateNewOrder(&newOrder)
	fmt.Println(Resp26.Msg)
	if !compareOrders(&newOrder, &Resp26.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}
	newOrder = Resp26.Data
	Resp27, _ := orderSvc.ReqUpdateOrder_OrderService(&newOrder)
	fmt.Println(Resp27.Msg)
	if !compareOrders(&newOrder, &Resp27.Data) {
		t.Errorf("【ReqDeleteOrder】unexpected returned order.")
		t.Skip()
	}
	Resp28, _ := orderSvc.ReqDeleteOrder_OrderService(Resp27.Data.Id)
	fmt.Println(Resp28.Msg)
}
