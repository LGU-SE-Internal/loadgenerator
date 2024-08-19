package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

// VerifyCodeBehaviorChain
func VerifyCode(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	verifyCode := generateVerifyCode()
	verifyCodeResp, err := cli.VerifyCode(verifyCode)
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if !verifyCodeResp {
		log.Errorf("Verification failed")
		return nil, err
	}

	ctx.Set(BooleanVerifyCode, verifyCodeResp)

	return nil, nil
}
