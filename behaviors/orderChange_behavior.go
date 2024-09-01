package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
)

const ()

var OrderChangeChain *Chain

func init() {
	OrderChangeChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(RefreshOrder, "RefreshOrder"),
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

	var rebookSvc service.
}
