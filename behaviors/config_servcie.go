package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

// ConfigBehaviorChain
func QueryConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype All Configs Test
	queryAllResp, err := cli.QueryAllConfigs()
	if err != nil {
		log.Errorf("QueryAllConfigs request failed, err %s", err)
		return nil, err
	}
	if queryAllResp.Status != 1 {
		log.Errorf("QueryAllConfigs status != 1")
		return nil, err
	}

	/*	Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`*/
	randomIndex := rand.Intn(len(queryAllResp.Data))
	ctx.Set(ConfigName, queryAllResp.Data[randomIndex].Name)
	ctx.Set(Value, queryAllResp.Data[randomIndex].Value)
	ctx.Set(Description, queryAllResp.Data[randomIndex].Description)

	return nil, nil
}
