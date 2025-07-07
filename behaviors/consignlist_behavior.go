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
		return nil, fmt.Errorf("ConsignList: unable to retrieve service client from context")
	}

	TheAccountId := ctx.Get(AccountID).(string)

	var consignSvc service.ConsignService = cli
	ConsignListResp, err := consignSvc.QueryByAccountId(TheAccountId)
	if err != nil {
		log.Errorf("ConsignList: QueryByAccountId failed for AccountID=%s, error=%v", TheAccountId, err)
		return nil, err
	}
	if ConsignListResp.Status != 1 {
		log.Warnf("ConsignList: QueryByAccountId returned non-success status for AccountID=%s, status=%v", TheAccountId, ConsignListResp.Status)
		return nil, fmt.Errorf("ConsignList: service returned status=%v for AccountID=%s", ConsignListResp.Status, TheAccountId)
	}
	log.Infof("ConsignList: QueryByAccountId succeeded for AccountID=%s, status=%v", TheAccountId, ConsignListResp.Status)

	return &(NodeResult{false}), nil // Chain End :D
}
