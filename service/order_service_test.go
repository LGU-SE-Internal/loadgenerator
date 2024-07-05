package service

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

func TestSvcImpl_ReqFindAllOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqFindAllOrder()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqCreateNewOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqCreateNewOrder(&Order{
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

func TestSvcImpl_ReqSaveOrderInfo(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqSaveOrderInfo(&Order{
		AccountId:              "test1",
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "790bcfd5-82d2-4717-aa9f-e00bef992268",
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

func TestSvcImpl_ReqAddCreateNewOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqAddCreateNewOrder(&Order{
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

func TestSvcImpl_ReqUpdateOrder_OrderService(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqUpdateOrder_OrderService(&Order{
		AccountId:              "test1",
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        RandomDecimalStringBetween(1, 10),
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "790bcfd5-82d2-4717-aa9f-e00bef992268",
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

func TestSvcImpl_ReqPayOrder(t *testing.T) {
	cli, _ := GetBasicClient()
	Resp, err := cli.ReqPayOrder("790bcfd5-82d2-4717-aa9f-e00bef992268")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetOrderPrice(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetOrderPrice("790bcfd5-82d2-4717-aa9f-e00bef992268")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqQueryOrders(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqQueryOrders(&Qi{
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

func TestSvcImpl_ReqQueryOrderForRefresh(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqQueryOrderForRefresh(&Qi{
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

func TestSvcImpl_ReqSecurityInfoCheck(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqSecurityInfoCheck(faker.Date(), "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqModifyOrder(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqModifyOrder("790bcfd5-82d2-4717-aa9f-e00bef992268", 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetTicketsList(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetTicketsList(&Seat{
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

func TestSvcImpl_ReqCalculateSoldTicket(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqCalculateSoldTicket(faker.Date(), GenerateTrainNumber())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetOrderById(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetOrderById("790bcfd5-82d2-4717-aa9f-e00bef992268")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqDeleteOrder_OrderService(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqDeleteOrder_OrderService("f72aa648-132e-43cd-8355-54819839deb9")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_End2End_OrderService(t *testing.T) {
	cli, _ := GetAdminClient()
	var orderSvc OrderService = cli

	Resp, err := orderSvc.ReqFindAllOrder()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	if len(Resp.Data) != 0 {
		t.Log("no data found.")
	}

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

	Resp0, err := orderSvc.ReqCreateNewOrder(&origin_order_0)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returned_order_0 := Resp0.Data
	rv_0 := reflect.ValueOf(&returned_order_0).Elem()
	ov_0 := reflect.ValueOf(&origin_order_0).Elem()

	for i := 0; i < rv_0.NumField(); i++ {

		fieldName := rv_0.Type().Field(i).Name
		if fieldName == "Id" {
			continue // 跳过 id 字段的比较
		}

		rf := rv_0.Field(i)
		of := ov_0.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv_0.Type().Field(i).Name, rf.Interface(), rv_0.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	Resp1, err := orderSvc.ReqGetOrderById(returned_order_0.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returned_order_1 := Resp1.Data
	rv_1 := reflect.ValueOf(&returned_order_1).Elem()
	ov_1 := reflect.ValueOf(&returned_order_0).Elem()

	for i := 0; i < rv_1.NumField(); i++ {

		rf := rv_1.Field(i)
		of := ov_1.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv_1.Type().Field(i).Name, rf.Interface(), rv_1.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	origin_order_1 := Order{
		AccountId:              uuid.NewString(),
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           faker.Name(),
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     returned_order_0.Id,
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            "G111",
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	Resp2, err := orderSvc.ReqSaveOrderInfo(&origin_order_1)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returned_order_2 := Resp2.Data
	rv_2 := reflect.ValueOf(&origin_order_1).Elem()
	ov_2 := reflect.ValueOf(&returned_order_2).Elem()

	for i := 0; i < rv_2.NumField(); i++ {

		rf := rv_2.Field(i)
		of := ov_2.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv_2.Type().Field(i).Name, rf.Interface(), rv_2.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	Resp3, err := orderSvc.ReqGetOrderById(returned_order_2.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}

	returned_order_3 := Resp3.Data
	rv_3 := reflect.ValueOf(&returned_order_3).Elem()
	ov_3 := reflect.ValueOf(&returned_order_2).Elem()

	for i := 0; i < rv_3.NumField(); i++ {

		rf := rv_3.Field(i)
		of := ov_3.Field(i)

		if !reflect.DeepEqual(rf.Interface(), of.Interface()) {
			t.Errorf("Request failed, returned_order.%s: %v, origin_order.%s: %v", rv_3.Type().Field(i).Name, rf.Interface(), rv_3.Type().Field(i).Name, of.Interface())
			t.Skip()
		}
	}

	Resp4, err := orderSvc.ReqGetOrderPrice(origin_order_1.Id)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
		t.Skip()
	}
	t.Log(Resp4)

	Resp5, err := orderSvc.ReqPayOrder(origin_order_1.Id)
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

	Resp7, err := orderSvc.ReqDeleteOrder_OrderService(origin_order_1.Id)
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
