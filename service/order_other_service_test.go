package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"strconv"
	"testing"
)

func TestSvcImpl_ReqFindAllOeder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqFindAllOeder_Other()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqCreateNewOeder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	AddResp, err := cli.ReqCreateNewOeder_Other(&Order{
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
	UpdateResp, err := cli.ReqSaveOrderInfo_Other(&Order{
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
	AddResp, err := cli.ReqAddCreateNewOrder_Other(&Order{
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
	UpdateResp, err := cli.ReqUpdateOrder_OrderService_Other(&Order{
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
	Resp, err := cli.ReqPayOrder_Other("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetOrderPrice_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetOrderPrice_Other("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqQueryOrders_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqQueryOrders_Other(&Qi{
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
	Resp, err := cli.ReqQueryOrderForRefresh_Other(&Qi{
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
	Resp, err := cli.ReqSecurityInfoCheck_Other(faker.Date(), "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqModifyOrder_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqModifyOrder_Other("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa", 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetTicketsList_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetTicketsList_Other(&Seat{
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
	Resp, err := cli.ReqCalculateSoldTicket_Other(faker.Date(), GenerateTrainNumber())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqGetOrderById_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqGetOrderById_Other("ee628cb0-6512-4dd0-ba1e-6eb5ccededaa")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_ReqDeleteOrder_OrderService_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, err := cli.ReqDeleteOrder_OrderService_Other("4a766f5d-ef7c-4629-aef4-04492e247503")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(Resp.Msg)
}

func TestSvcImpl_End2End_OrderService_Other(t *testing.T) {
	cli, _ := GetAdminClient()
	Resp, _ := cli.ReqCreateNewOeder_Other(&Order{
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
	Resp, _ = cli.ReqSaveOrderInfo_Other(&originOrder)
	fmt.Println(Resp.Msg)
	ArrResp, _ := cli.ReqGetOrderPrice_Other(originOrder.Id)
	fmt.Println(ArrResp.Msg)
	Resp, _ = cli.ReqPayOrder_Other(originOrder.Id)
	fmt.Println(Resp.Msg)
	DataResp, _ := cli.ReqQueryOrders_Other(&Qi{
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
	DataResp, _ = cli.ReqQueryOrderForRefresh_Other(&Qi{
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
	Resp, _ = cli.ReqSecurityInfoCheck_Other(originOrder.BoughtDate, originOrder.AccountId)
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqModifyOrder_Other(originOrder.Id, 0)
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqGetTicketsList_Other(&Seat{
		DestStation:  RandomProvincialCapitalEN(),
		SeatType:     2,
		StartStation: RandomProvincialCapitalEN(),
		Stations:     nil,
		TotalNum:     0,
		TrainNumber:  GenerateTrainNumber(),
		TravelDate:   faker.Date(),
	})
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqCalculateSoldTicket_Other(faker.Date(), GenerateTrainNumber())
	fmt.Println(Resp.Msg)
	Resp, _ = cli.ReqGetOrderById_Other(originOrder.Id)
	fmt.Println(Resp.Msg)
	StringResp, _ := cli.ReqDeleteOrder_OrderService_Other(originOrder.Id)
	fmt.Println(StringResp.Msg)
	Resp, _ = cli.ReqAddCreateNewOrder_Other(&Order{
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
	Resp, _ = cli.ReqUpdateOrder_OrderService_Other(&Resp.Data)
	fmt.Println(Resp.Msg)
}
