package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const ()

var OrderCancelChain *Chain

func init() {
	OrderCancelChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrder, "RefreshOrder"),
		NewFuncNode(OrderCalculate, "OrderCalculate"),
		NewFuncNode(OrderCancel, "OrderCancel"),
	)

	fmt.Println(OrderCancelChain.VisualizeChain(0))
}

func OrderCalculate(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheOrderId := ctx.Get(OrderId).(string)

	var cancelSvc service.CancelService = cli
	OrderCalculateResp, err := cancelSvc.ReqCalculate(TheOrderId)
	if err != nil {
		return nil, err
	}
	if OrderCalculateResp.Status != 1 {
		return nil, fmt.Errorf("OrderCalculate fail. Preserve Status is %v", OrderCalculateResp.Status)
	}

	log.Infof("OrderCalculate success. Preserve Status is %v", OrderCalculateResp.Status)

	return nil, nil
}

func OrderCancel(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheOrderId := ctx.Get(OrderId).(string)
	TheAccountId := ctx.Get(AccountID).(string)

	var cancelSvc service.CancelService = cli
	OrderCancelResp, err := cancelSvc.ReqCancelTicket(TheOrderId, TheAccountId)
	if err != nil {
		return nil, err
	}
	if OrderCancelResp.Status != 1 {
		return nil, fmt.Errorf("OrderCancel fail. Preserve Status is %v", OrderCancelResp.Status)
	}

	log.Infof("[Success]OrderCancel success! Preserve Status is %v", OrderCancelResp.Status)

	return &(NodeResult{false}), nil // Chain End :D
}
