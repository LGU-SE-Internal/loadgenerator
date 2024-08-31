package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
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

	fmt.Println(OrderConsignChain.VisualizeChain(0))
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
		return nil, err
	}
	if OrderQueryResp.Status != 1 {
		return nil, fmt.Errorf("OrderQuery Fails. Status != 1. Order Query status is %d", OrderQueryResp.Status)
	}
	log.Infof("[Success]The Status is: %v, and OrderConsignResp Data: %v", OrderQueryResp.Status, OrderQueryResp.Data)

	return &(NodeResult{false}), nil // Chain End :D
}

func OrderConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheConsign := service.Consign{}

	var consignSvc service.ConsignService = cli
	OrderConsignResp, err := consignSvc.UpdateConsignRecord(&TheConsign)
	if err != nil {
		return nil, err
	}
	if OrderConsignResp.Status != 1 {
		return nil, fmt.Errorf("OrderConsign Fails. Status != 1. Order consign status is %d", OrderConsignResp.Status)
	}
	log.Infof("[Success]The Status is: %v, and OrderConsignResp Data: %v", OrderConsignResp.Status, OrderConsignResp.Data)

	return &(NodeResult{true}), nil // Chain End :D
}
