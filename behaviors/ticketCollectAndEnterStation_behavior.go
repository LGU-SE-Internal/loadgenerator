package behaviors

import (
	"fmt"

	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const ()

var TicketCollectAndEnterStationChain *Chain

func init() {
	TicketCollectAndEnterStationChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrderOther, "RefreshOrderOther"),
		NewFuncNode(TicketCollect, "TicketCollect"),
		NewFuncNode(EnterStation, "EnterStation"),
	)
	// fmt.Println(TicketCollectAndEnterStationChain.VisualizeChain(0))
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
	log.Infof("[Success]The Status is: %v, and TicketCollectResp Msg: %v", TicketCollectResp.Status, TicketCollectResp.Msg)

	//return &(NodeResult{false}), nil // Chain End :D
	return nil, nil
}

func EnterStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheOrderId := ctx.Get(OrderId).(string)

	var executeSvc service.ExecuteService = cli
	EnterStationResp, err := executeSvc.ReqExecuteTicket(TheOrderId)
	if err != nil {
		return nil, err
	}
	if EnterStationResp.Status != 1 {
		return nil, fmt.Errorf("execute Ticket Failed. EnterStationResp Status != 1, get: %v", EnterStationResp.Status)
	}
	log.Infof("[Success] The EnterStationResp Status: %v, and EnterStationResp Msg is: %v", EnterStationResp.Status, EnterStationResp.Msg)

	return &(NodeResult{false}), nil // Chain End :D
}
