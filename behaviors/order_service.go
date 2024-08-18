package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Resp, err := cli.ReqFindAllOrder()
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if len(Resp.Data) == 0 {
		log.Errorf("no data found.")
		return nil, err
	}

	randomIndex := rand.Intn(len(Resp.Data))
	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, Resp.Data[randomIndex].BoughtDate)
	ctx.Set(CoachNumber, Resp.Data[randomIndex].CoachNumber)
	ctx.Set(ContactsDocumentNumber, Resp.Data[randomIndex].ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, Resp.Data[randomIndex].ContactsName)
	ctx.Set(DifferenceMoney, Resp.Data[randomIndex].DifferenceMoney)
	ctx.Set(SeatClass, Resp.Data[randomIndex].SeatClass)
	ctx.Set(SeatNumber, Resp.Data[randomIndex].SeatNumber)
	ctx.Set(Status, Resp.Data[randomIndex].Status)
	ctx.Set(TrainNumber, Resp.Data[randomIndex].TrainNumber)
	ctx.Set(TravelDate, Resp.Data[randomIndex].TravelDate)
	ctx.Set(TravelTime, Resp.Data[randomIndex].TravelTime)

	return nil, nil
}

func CreateOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	originOrder0 := service.Order{
		AccountId:              ctx.Get(AccountID).(string),
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		//ContactsName:           ctx.Get(ContactsName).(string),
		ContactsName:    ctx.Get(Name).(string),
		DifferenceMoney: "",
		DocumentType:    0,
		//From:                   ctx.Get(From).(string),
		From: ctx.Get(From).(string), // First, create/query get the station;
		// then put them here -> If you want to create a new Order, you have to do the whole process.
		Id:         "nil",
		Price:      RandomDecimalStringBetween(1, 10),
		SeatClass:  GetTrainTicketClass(),
		SeatNumber: service.GenerateSeatNumber(),
		Status:     0,
		//To:                     ctx.Get(To).(string),
		To:          ctx.Get(To).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		TravelDate:  getRandomTime(),
		TravelTime:  generateRandomTime(),
	}

	CreateNewOrderResp, err := cli.ReqCreateNewOrder(&originOrder0)
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if CreateNewOrderResp.Status != 1 {
		log.Errorf("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, CreateNewOrderResp.Data.BoughtDate)
	ctx.Set(CoachNumber, CreateNewOrderResp.Data.CoachNumber)
	ctx.Set(ContactsDocumentNumber, CreateNewOrderResp.Data.ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, CreateNewOrderResp.Data.ContactsName)
	ctx.Set(DifferenceMoney, CreateNewOrderResp.Data.DifferenceMoney)
	ctx.Set(SeatClass, CreateNewOrderResp.Data.SeatClass)
	ctx.Set(SeatNumber, CreateNewOrderResp.Data.SeatNumber)
	ctx.Set(Status, CreateNewOrderResp.Data.Status)
	ctx.Set(TrainNumber, CreateNewOrderResp.Data.TrainNumber)
	ctx.Set(TravelDate, CreateNewOrderResp.Data.TravelDate)
	ctx.Set(TravelTime, CreateNewOrderResp.Data.TravelTime)

	return nil, nil
}
