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

	// fmt.Println(OrderCancelChain.VisualizeChain(0))
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
		log.Errorf("OrderCalculate failed: error occurred while requesting calculation for order_id=%s: %v", TheOrderId, err)
		return nil, err
	}
	if OrderCalculateResp.Status != 1 {
		log.Errorf("OrderCalculate failed: unexpected response status for order_id=%s, status=%v", TheOrderId, OrderCalculateResp.Status)
		return nil, fmt.Errorf("OrderCalculate failed, response status: %v", OrderCalculateResp.Status)
	}

	log.Infof("OrderCalculate succeeded: order_id=%s, response status=%v", TheOrderId, OrderCalculateResp.Status)

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
		log.Errorf("OrderCancel failed: error occurred while requesting cancellation for order_id=%s, account_id=%s: %v", TheOrderId, TheAccountId, err)
		return nil, err
	}
	if OrderCancelResp.Status != 1 {
		log.Errorf("OrderCancel failed: unexpected response status for order_id=%s, account_id=%s, status=%v", TheOrderId, TheAccountId, OrderCancelResp.Status)
		return nil, fmt.Errorf("OrderCancel failed, response status: %v", OrderCancelResp.Status)
	}

	log.Infof("OrderCancel succeeded: order_id=%s, account_id=%s, response status=%v", TheOrderId, TheAccountId, OrderCancelResp.Status)

	return &(NodeResult{false}), nil // Chain End :D
}
