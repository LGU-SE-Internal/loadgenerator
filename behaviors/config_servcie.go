package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

// ConfigBehaviorChain
func QueryConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		log.Errorf("QueryConfig: failed to get service client from context, got type: %T", ctx.Get(Client))
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype All Configs Test
	queryAllResp, err := cli.QueryAllConfigs()
	if err != nil {
		log.Errorf("QueryConfig: QueryAllConfigs request failed, error: %v", err)
		return nil, err
	}
	if queryAllResp.Status != 1 {
		log.Errorf("QueryConfig: QueryAllConfigs returned unexpected status: %d, expected: 1", queryAllResp.Status)
		return nil, fmt.Errorf("unexpected status from QueryAllConfigs: %d", queryAllResp.Status)
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
