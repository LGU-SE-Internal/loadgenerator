package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const ()

var TicketCollectChain *Chain

func init() {
	TicketCollectChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrderOther, "RefreshOrderOther"),
		NewFuncNode(TicketCollect, "TicketCollect"),
	)
	fmt.Println(TicketCollectChain.VisualizeChain(0))
}

func TicketCollect(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheOrderId := ctx.Get(OrderId).(string)

	var executeSvc service.ExecuteService = cli
	TicketCollectResp, err := executeSvc.ReqCollectTicket(TheOrderId)
	if err != nil {
		return nil, err
	}
	if TicketCollectResp.Status != 1 {
		return nil, fmt.Errorf("collect tickets fail. TicketCollectResp.Status != 1, get %v", TicketCollectResp.Status)
	}
	log.Infof("[Success]The Status is: %v, and TicketCollectResp Data: %v", TicketCollectResp.Status, TicketCollectResp.Data)

	return &(NodeResult{false}), nil // Chain End :D
}
