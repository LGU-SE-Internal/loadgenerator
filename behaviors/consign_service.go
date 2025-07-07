package behaviors

import (
	"fmt"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

// ConsignBehaviorsChain
func QueryConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		log.Error("Failed to retrieve service client from context: type assertion to *service.SvcImpl failed")
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype consign records by order ID
	TheOrderId := ctx.Get(OrderId).(string)
	consignsByOrderId, err := cli.QueryByOrderId(TheOrderId)
	if err != nil {
		log.Errorf("Consign query failed for orderId '%s': %v", TheOrderId, err)
		return nil, err
	}
	if consignsByOrderId.Status != 1 {
		log.Errorf("Consign query for orderId '%s' returned unexpected status %d. Full response: %+v", TheOrderId, consignsByOrderId.Status, consignsByOrderId)
		return nil, fmt.Errorf("unexpected status from QueryByOrderId: %d", consignsByOrderId.Status)
	}

	ctx.Set(ID, consignsByOrderId.Data.Id)
	ctx.Set(OrderId, consignsByOrderId.Data.OrderId)
	//ctx.Set(AccountID, consignsByOrderId.Data.AccountId)
	ctx.Set(HandleDate, consignsByOrderId.Data.HandleDate)
	ctx.Set(TargetDate, consignsByOrderId.Data.TargetDate)
	ctx.Set(From, consignsByOrderId.Data.From)
	ctx.Set(To, consignsByOrderId.Data.To)
	ctx.Set(Consignee, consignsByOrderId.Data.Consignee)
	ctx.Set(Phone, consignsByOrderId.Data.Phone)
	ctx.Set(Weight, consignsByOrderId.Data.Weight)
	ctx.Set(Price, consignsByOrderId.Data.Price)

	return nil, nil
}

func CreateConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		log.Error("Failed to retrieve service client from context: type assertion to *service.SvcImpl failed")
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedId := faker.UUIDHyphenated()
	MockedAccountId := ctx.Get(AccountID).(string)
	MockedOrderId := ctx.Get(OrderId).(string)
	MockedHandleDate := ctx.Get(DepartureTime).(string)
	//MockedHandleDate := ""
	MockedTargetDate := ctx.Get(DepartureTime).(string)
	//MockedTargetDate := ""
	MockedFromPlace := ctx.Get(StartStation).(string)
	MockedToPlace := ctx.Get(EndStation).(string)
	MockedConsignee := faker.Name()
	MockedPhone := faker.Phonenumber()
	MockedWeight := GenerateWeight()

	// Insert a new consign record
	insertReq := service.Consign{
		ID:         MockedId,
		OrderID:    MockedOrderId,
		AccountID:  MockedAccountId,
		HandleDate: MockedHandleDate,
		TargetDate: MockedTargetDate,
		From:       MockedFromPlace,
		To:         MockedToPlace,
		Consignee:  MockedConsignee,
		Phone:      MockedPhone,
		Weight:     MockedWeight,
		IsWithin:   BooleanIsWithin(MockedWeight),
	}
	insertResp, err := cli.InsertConsignRecord(&insertReq)
	if err != nil {
		log.Errorf("Failed to insert consign record. Request: %+v, Error: %v", insertReq, err)
		return nil, err
	}
	if insertResp.Msg == "Already exists" {
		log.Warnf("Consign record already exists. Request: %+v", insertReq)
		return nil, fmt.Errorf("Consign already exists")
	}
	if insertResp.Status != 1 {
		log.Errorf("InsertConsignRecord returned non-success status %d. Request: %+v, Response: %+v", insertResp.Status, insertReq, insertResp)
		return nil, fmt.Errorf("InsertConsignRecord returned unexpected status: %d", insertResp.Status)
	}
	isMatch := false
	if /*insertResp.Data.ID == insertReq.ID &&*/
	/*insertResp.Data.IsWithin == insertReq.IsWithin &&*/
	insertResp.Data.AccountID == insertReq.AccountID &&
		insertResp.Data.From == insertReq.From &&
		insertResp.Data.Consignee == insertReq.Consignee &&
		insertResp.Data.OrderID == insertReq.OrderID &&
		insertResp.Data.Phone == insertReq.Phone &&
		insertResp.Data.TargetDate == insertReq.TargetDate &&
		insertResp.Data.HandleDate == insertReq.HandleDate &&
		insertResp.Data.To == insertReq.To &&
		insertResp.Data.Weight == insertReq.Weight {
		isMatch = true
	}
	if !isMatch {
		log.Errorf("Mismatch between consign creation request and response. Expected: %+v, Actual: %+v", insertReq, insertResp.Data)
		return nil, fmt.Errorf("consign creation result mismatch")
	}
	//log.Errorf("InsertConsignRecord response: %+v", insertResp)
	//existedConsign := insertResp.Data

	ctx.Set(ID, insertResp.Data.ID)
	//ctx.Set(OrderID, insertResp.Data.OrderID)
	//ctx.Set(AccountID, insertResp.Data.AccountID)
	ctx.Set(HandleDate, insertResp.Data.HandleDate)
	ctx.Set(TargetDate, insertResp.Data.TargetDate)
	ctx.Set(From, insertResp.Data.From)
	ctx.Set(To, insertResp.Data.To)
	ctx.Set(Consignee, insertResp.Data.Consignee)
	ctx.Set(Phone, insertResp.Data.Phone)
	ctx.Set(Weight, insertResp.Data.Weight)
	ctx.Set(IsWithin, insertResp.Data.IsWithin)

	return nil, nil
}
