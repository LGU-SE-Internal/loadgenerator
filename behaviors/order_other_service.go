package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	GetResp, err := cli.ReqFindAllOrderOther()

	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if GetResp.Status != 1 {
		log.Errorf("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetResp.Data))
	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, GetResp.Data[randomIndex].BoughtDate)
	ctx.Set(CoachNumber, GetResp.Data[randomIndex].CoachNumber)
	ctx.Set(ContactsDocumentNumber, GetResp.Data[randomIndex].ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, GetResp.Data[randomIndex].ContactsName)
	ctx.Set(DifferenceMoney, GetResp.Data[randomIndex].DifferenceMoney)
	ctx.Set(SeatClass, GetResp.Data[randomIndex].SeatClass)
	ctx.Set(SeatNumber, GetResp.Data[randomIndex].SeatNumber)
	ctx.Set(Status, GetResp.Data[randomIndex].Status)
	ctx.Set(TrainNumber, GetResp.Data[randomIndex].TrainNumber)
	ctx.Set(TravelDate, GetResp.Data[randomIndex].TravelDate)
	ctx.Set(TravelTime, GetResp.Data[randomIndex].TravelTime)

	return nil, nil
}

func CreateOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AddResp, err := cli.ReqCreateNewOrderOther(&service.Order{
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
	})

	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if AddResp.Status != 1 {
		log.Errorf("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	ctx.Set(BoughtDate, AddResp.Data.BoughtDate)
	ctx.Set(CoachNumber, AddResp.Data.CoachNumber)
	ctx.Set(ContactsDocumentNumber, AddResp.Data.ContactsDocumentNumber)
	ctx.Set(Name, AddResp.Data.ContactsName)
	ctx.Set(DifferenceMoney, AddResp.Data.DifferenceMoney)
	ctx.Set(SeatClass, AddResp.Data.SeatClass)
	ctx.Set(SeatNumber, AddResp.Data.SeatNumber)
	ctx.Set(Status, AddResp.Data.Status)
	ctx.Set(TrainNumber, AddResp.Data.TrainNumber)
	ctx.Set(TravelDate, AddResp.Data.TravelDate)
	ctx.Set(TravelTime, AddResp.Data.TravelTime)

	return nil, nil
}
