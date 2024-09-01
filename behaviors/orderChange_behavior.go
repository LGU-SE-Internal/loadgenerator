package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const (
	OldTripID = "oldTripID"
)

var OrderChangeChain *Chain

func init() {
	OrderChangeChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrder, "RefreshOrder"),
		NewFuncNode(QueryTrain, "QueryTrain"),
		NewFuncNode(ChooseRoute, "ChooseRoute"),
		NewFuncNode(QueryTripInfo, "QueryTripInfo"),
		NewFuncNode(OrderRebook, "OrderRebook"),
	)

	fmt.Println(OrderChangeChain.VisualizeChain(0))
}

func OrderRebook(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	RebookInfo := service.RebookInfo{
		LoginID:   ctx.Get(AccountID).(string),
		OrderID:   ctx.Get(OrderId).(string),
		OldTripID: ctx.Get(OldTripID).(string),
		TripID:    ctx.Get(TripID).(string),
		SeatType:  ctx.Get(SeatClass).(int),
		Date:      ctx.Get(DepartureTime).(string),
	}

	var rebookSvc service.ReBookService = cli
	OrderRebookResp, err := rebookSvc.Rebook(&RebookInfo)
	if err != nil {
		return nil, err
	}
	if OrderRebookResp.Status != 1 {
		return nil, fmt.Errorf("order rebook fail. OrderRebookResp Status is %d instead of 1", OrderRebookResp.Status)
	}
	log.Infof("[Success!] OrderRebookResp Status is %d", OrderRebookResp.Status)

	return &(NodeResult{false}), nil // Chain End :D
}
