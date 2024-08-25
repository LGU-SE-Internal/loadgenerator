package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const ()

var NormalOrderPayChain *Chain

func init() {
	NormalOrderPayChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrder, "RefreshOrder"),
		NewFuncNode(OrderPay, "OrderPay"),
	)
	fmt.Println(NormalOrderPayChain.VisualizeChain(0))
}

func OrderPay(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	tripPaymentInfo := service.TripPayment{
		TripId:  ctx.Get(TrainNumber).(string), // TrainNumber = TripID
		OrderId: ctx.Get(OrderId).(string),
		Price:   ctx.Get(Price).(string),
		UserId:  ctx.Get(UserId).(string),
	}

	var insidePaymentSvc service.InsidePaymentService = cli
	InsidePaymentResp, err := insidePaymentSvc.ReqPay_InsidePayment(&tripPaymentInfo)
	if err != nil {
		return nil, err
	}
	if InsidePaymentResp.Status != 1 {
		return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", InsidePaymentResp.Status)
	}
	log.Infof("The Status is: %v, and PreserveResp Data: %v", InsidePaymentResp.Status, InsidePaymentResp.Data)

	return &(NodeResult{false}), nil // Chain End :D
}
