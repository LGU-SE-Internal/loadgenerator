package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

// Admin Order Behavior Chain - 管理订单的增删改查操作
var AdminOrderChain *Chain

func init() {
	AdminOrderChain = NewChain(
		NewFuncNode(LoginAdmin, "LoginAdmin"),
		NewFuncNode(AdminQueryAllOrders, "AdminQueryAllOrders"),
	)

	// 添加后续的订单管理操作链
	AdminOrderChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddOrder, "AdminAddOrder"),
		NewFuncNode(AdminUpdateOrder, "AdminUpdateOrder"),
	), 0.5)

	AdminOrderChain.AddNextChain(NewChain(
		NewFuncNode(AdminDeleteOrder, "AdminDeleteOrder"),
	), 0.5)
}

// AdminQueryAllOrders 查询所有订单
func AdminQueryAllOrders(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.ReqGetAllOrders()
	if err != nil {
		log.Errorf("AdminQueryAllOrders failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminQueryAllOrders returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	if len(resp.Data) > 0 {
		randomIndex := rand.Intn(len(resp.Data))
		ctx.Set(OrderId, resp.Data[randomIndex].Id)
		ctx.Set(AccountID, resp.Data[randomIndex].AccountId)
		ctx.Set(TrainNumber, resp.Data[randomIndex].TrainNumber)
		ctx.Set(From, resp.Data[randomIndex].From)
		ctx.Set(To, resp.Data[randomIndex].To)
		ctx.Set(SeatClass, resp.Data[randomIndex].SeatClass)
		ctx.Set(Status, resp.Data[randomIndex].Status)
		ctx.Set(TravelDate, resp.Data[randomIndex].TravelDate)
	}

	log.Infof("AdminQueryAllOrders returned %d orders", len(resp.Data))
	return nil, nil
}

// AdminAddOrder 添加订单
func AdminAddOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	order := &service.Order{
		AccountId:              faker.UUIDHyphenated(),
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		ContactsName:           faker.Name(),
		DifferenceMoney:        "",
		DocumentType:           rand.Intn(2),
		From:                   generateRandomCityName(),
		Id:                     "",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             service.GenerateSeatNumber(),
		Status:                 0,
		To:                     generateRandomCityName(),
		TrainNumber:            GenerateTripId(),
		TravelDate:             getRandomTime(),
		TravelTime:             generateRandomTime(),
	}

	resp, err := cli.ReqAddOrder(order)
	if err != nil {
		log.Errorf("AdminAddOrder failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddOrder returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(OrderId, resp.Data.Id)
	ctx.Set(AccountID, resp.Data.AccountId)
	ctx.Set(TrainNumber, resp.Data.TrainNumber)
	ctx.Set(From, resp.Data.From)
	ctx.Set(To, resp.Data.To)

	log.Infof("AdminAddOrder success: orderId=%s", resp.Data.Id)
	return nil, nil
}

// AdminUpdateOrder 更新订单
func AdminUpdateOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	orderId, ok := ctx.Get(OrderId).(string)
	if !ok || orderId == "" {
		log.Warn("No order ID found in context, skipping update")
		return nil, nil
	}

	accountId, _ := ctx.Get(AccountID).(string)
	trainNumber, _ := ctx.Get(TrainNumber).(string)

	order := &service.Order{
		Id:                     orderId,
		AccountId:              accountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		ContactsName:           faker.Name(),
		DifferenceMoney:        "",
		DocumentType:           rand.Intn(2),
		From:                   generateRandomCityName(),
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             service.GenerateSeatNumber(),
		Status:                 rand.Intn(6), // 0-5 different status
		To:                     generateRandomCityName(),
		TrainNumber:            trainNumber,
		TravelDate:             getRandomTime(),
		TravelTime:             generateRandomTime(),
	}

	resp, err := cli.ReqUpdateOrder(order)
	if err != nil {
		log.Errorf("AdminUpdateOrder failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminUpdateOrder returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminUpdateOrder success: orderId=%s", orderId)
	return nil, nil
}

// AdminDeleteOrder 删除订单
func AdminDeleteOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	orderId, ok := ctx.Get(OrderId).(string)
	if !ok || orderId == "" {
		log.Warn("No order ID found in context, skipping delete")
		return nil, nil
	}

	trainNumber, ok := ctx.Get(TrainNumber).(string)
	if !ok || trainNumber == "" {
		trainNumber = GenerateTripId()
	}

	resp, err := cli.ReqDeleteOrder(orderId, trainNumber)
	if err != nil {
		log.Errorf("AdminDeleteOrder failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminDeleteOrder returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminDeleteOrder success: orderId=%s", orderId)
	return nil, nil
}
