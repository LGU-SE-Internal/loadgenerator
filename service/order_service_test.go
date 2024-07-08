package service

import (
	"fmt"
	"reflect"
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

	originOrder0 := Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
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
	rv0 := reflect.ValueOf(&returnedOrder0).Elem()
	ov0 := reflect.ValueOf(&originOrder0).Elem()

	for i := 0; i < rv0.NumField(); i++ {

		fieldName := rv0.Type().Field(i).Name
		if fieldName == "Id" {
			continue // 跳过 id 字段的比较
		}

		rf := rv0.Field(i)
		of := ov0.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv0.Type().Field(i).Name, rf.Interface(), rv0.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	Resp1, err := orderSvc.ReqGetOrderById(returnedOrder0.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder1 := Resp1.Data
	rv1 := reflect.ValueOf(&returnedOrder1).Elem()
	ov1 := reflect.ValueOf(&returnedOrder0).Elem()

	for i := 0; i < rv1.NumField(); i++ {

		rf := rv1.Field(i)
		of := ov1.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv1.Type().Field(i).Name, rf.Interface(), rv1.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	originOrder1 := Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
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
	rv2 := reflect.ValueOf(&originOrder1).Elem()
	ov2 := reflect.ValueOf(&returnedOrder2).Elem()

	for i := 0; i < rv2.NumField(); i++ {

		rf := rv2.Field(i)
		of := ov2.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv2.Type().Field(i).Name, rf.Interface(), rv2.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	Resp3, err := orderSvc.ReqGetOrderById(returnedOrder2.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returnedOrder3 := Resp3.Data
	rv3 := reflect.ValueOf(&returnedOrder3).Elem()
	ov3 := reflect.ValueOf(&returnedOrder2).Elem()

	for i := 0; i < rv3.NumField(); i++ {

		rf := rv3.Field(i)
		of := ov3.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv3.Type().Field(i).Name, rf.Interface(), rv3.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
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

	Resp6, err := orderSvc.ReqQueryOrders(&Qi{
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
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp6)

	Resp7, err := orderSvc.ReqDeleteOrder_OrderService(originOrder1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp7)
}

func TestSvcImpl_End2End_OrderService_v2(t *testing.T) {
	cli, _ := GetAdminClient()

	Resp6, err := cli.ReqQueryOrders(&Qi{
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
		t.Errorf("Request failed, err %s", err)
	}
	t.Log(Resp6)

	origin_order_0 := Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
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

	Resp0, err := cli.ReqCreateNewOrder(&origin_order_0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}

	originOrder := Resp0.Data

	DataResp, _ := cli.ReqQueryOrderForRefresh(&Qi{
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
	fmt.Println(DataResp.Msg)
	Resp, _ := cli.ReqSecurityInfoCheck(originOrder.BoughtDate, originOrder.AccountId)
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqModifyOrder(originOrder.Id, 0)
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqGetTicketsList(&Seat{
		DestStation:  RandomProvincialCapitalEN(),
		SeatType:     2,
		StartStation: RandomProvincialCapitalEN(),
		Stations:     nil,
		TotalNum:     0,
		TrainNumber:  GenerateTrainNumber(),
		TravelDate:   faker.Date(),
	})
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqCalculateSoldTicket(faker.Date(), GenerateTrainNumber())
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqGetOrderById(originOrder.Id)
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqDeleteOrder_OrderService(originOrder.Id)
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqAddCreateNewOrder(&Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
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
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqUpdateOrder_OrderService(&Resp.Data)
	fmt.Println(Resp.Msg)
}
