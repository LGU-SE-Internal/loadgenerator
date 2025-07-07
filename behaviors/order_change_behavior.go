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
		NewFuncNode(RefreshOrderOther, "RefreshOrderOther"),
		NewFuncNode(ChooseRoute, "ChooseRoute"),
		NewFuncNode(QueryTrain, "QueryTrain"),
		NewFuncNode(QueryTripInfo, "QueryTripInfo"),
		NewFuncNode(QuerySeatInfo, "QuerySeatInfo"),
		NewFuncNode(OrderRebook, "OrderRebook"),
	)

	// fmt.Println(OrderChangeChain.VisualizeChain(0))
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
		log.Errorf("Order rebooking failed: %v", err)
		return nil, err
	}
	if OrderRebookResp.Status == 0 || OrderRebookResp.Msg == "You can only change the ticket before the train start or within 2 hours after the train start." {
		log.Warnf("Order rebooking not allowed: %s", OrderRebookResp.Msg)
		return &(NodeResult{false}), nil // Chain End
	} else if OrderRebookResp.Status != 1 {
		log.Errorf("Order rebooking failed: unexpected status %d, message: %s", OrderRebookResp.Status, OrderRebookResp.Msg)
		return nil, fmt.Errorf("OrderRebookResp failed: status=%d, msg=%v", OrderRebookResp.Status, OrderRebookResp.Msg)
	}

	log.Infof("Order rebooking succeeded: status=%d, orderID=%s, newTripID=%s", OrderRebookResp.Status, RebookInfo.OrderID, RebookInfo.TripID)

	return &(NodeResult{false}), nil // Chain End
}
