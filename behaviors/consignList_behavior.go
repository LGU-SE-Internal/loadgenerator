package behaviors

import (
	"fmt"

	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const ()

var ConsignListChain *Chain

func init() {
	ConsignListChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(ConsignList, "ConsignList"),
	)
	// fmt.Println(ConsignListChain.VisualizeChain(0))
}

func ConsignList(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheAccountId := ctx.Get(AccountID).(string)

	var consignSvc service.ConsignService = cli
	ConsignListResp, err := consignSvc.QueryByAccountId(TheAccountId)
	if err != nil {
		return nil, err
	}
	if ConsignListResp.Status != 1 {
		return nil, fmt.Errorf("consign service Status != 1, accountId: %s", TheAccountId)
	}
	log.Infof("[Success] get consign list success %v", ConsignListResp.Status)

	return &(NodeResult{false}), nil // Chain End :D
}
