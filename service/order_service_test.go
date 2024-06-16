package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"strconv"
	"testing"
)

func TestSvcImpl_ReqFindAllOeder(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqFindAllOeder()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqCreateNewOeder(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqCreateNewOeder(&Order{
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
	Resp, _ := cli.ReqCreateNewOeder(&Order{
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
	fmt.Println(Resp.Msg)
	originOrder := Resp.Data
	Resp, _ = cli.ReqSaveOrderInfo(&originOrder)
	fmt.Println(Resp.Msg)
	ArrResp, _ := cli.ReqGetOrderPrice(originOrder.Id)
	fmt.Println(ArrResp.Msg)
	Resp, _ = cli.ReqPayOrder(originOrder.Id)
	fmt.Println(Resp.Msg)
	DataResp, _ := cli.ReqQueryOrders(&Qi{
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
	DataResp, _ = cli.ReqQueryOrderForRefresh(&Qi{
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
	Resp, _ = cli.ReqSecurityInfoCheck(originOrder.BoughtDate, originOrder.AccountId)
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
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqUpdateOrder_OrderService(&Resp.Data)
	fmt.Println(Resp.Msg)
}
