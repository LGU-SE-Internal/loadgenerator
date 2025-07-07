package behaviors

import (
	"fmt"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

const ()

var OrderConsignBehaviorChain *Chain
var OrderConsignChain *Chain

func init() {
	OrderConsignChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrder, "RefreshOrder"),
		NewFuncNode(OrderQuery, "OrderQuery"),
		NewFuncNode(OrderConsign, "OrderConsign"),
	)

	// fmt.Println(OrderConsignChain.VisualizeChain(0))
}

func OrderQuery(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheOrderId := ctx.Get(OrderId).(string)

	var consignSvc service.ConsignService = cli
	OrderQueryResp, err := consignSvc.QueryByOrderId(TheOrderId)
	if err != nil {
		log.Errorf("Failed to query order by ID [%s]: %v", TheOrderId, err)
		return nil, err
	}
	if OrderQueryResp.Status == 0 {
		log.Infof("Order [%s] has not been consigned. Proceeding to consign.", TheOrderId)
		TheConsignee := faker.Name()
		ThePhone := faker.Phonenumber()
		TheWeight := GenerateWeight()
		ctx.Set(ID, OrderQueryResp.Data.Id)
		ctx.Set(Consignee, TheConsignee)
		ctx.Set(Phone, ThePhone)
		ctx.Set(Weight, TheWeight)
		log.Infof("Order [%s] is ready for consignment. Consignee: %s, Phone: %s, Weight: %.2f", TheOrderId, TheConsignee, ThePhone, TheWeight)
		return nil, nil
	} else {
		log.Warnf("Order [%s] has already been consigned. Skipping consignment.", TheOrderId)
		return &(NodeResult{false}), nil // Chain End
	}
}

func OrderConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheConsign := service.Consign{
		ID:         ctx.Get(ID).(string),
		OrderID:    ctx.Get(OrderId).(string),
		AccountID:  ctx.Get(AccountID).(string),
		HandleDate: ctx.Get(HandleDate).(string),
		TargetDate: ctx.Get(TargetDate).(string),
		From:       ctx.Get(From).(string),
		To:         ctx.Get(To).(string),
		Consignee:  ctx.Get(Consignee).(string),
		Phone:      ctx.Get(Phone).(string),
		Weight:     ctx.Get(Weight).(float64),
		IsWithin:   BooleanIsWithin(ctx.Get(Weight).(float64)),
	}

	var consignSvc service.ConsignService = cli
	OrderConsignResp, err := consignSvc.UpdateConsignRecord(&TheConsign)
	if err != nil {
		log.Errorf("Failed to update consignment record for OrderID [%s]: %v", TheConsign.OrderID, err)
		return nil, err
	}
	if OrderConsignResp.Status != 1 {
		log.Errorf("Consignment failed for OrderID [%s]. Unexpected status: %d", TheConsign.OrderID, OrderConsignResp.Status)
		return nil, fmt.Errorf("OrderConsign Fails. Status != 1. Order consign status is %d", OrderConsignResp.Status)
	}
	log.Infof("Consignment successful for OrderID [%s]. Status: %d, Response Data: %+v", TheConsign.OrderID, OrderConsignResp.Status, OrderConsignResp.Data)

	return &(NodeResult{true}), nil // Chain End
}
