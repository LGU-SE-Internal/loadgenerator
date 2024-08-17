package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QuerySecurity(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Get All Security Configs
	configs, err3 := cli.FindAllSecurityConfig()
	if err3 != nil {
		log.Errorf("FindAllSecurityConfig failed: %v", err3)
		return nil, err3
	}
	if configs.Status != 1 {
		log.Errorf("[Security Service]Status != 1")
		return nil, err3
	}

	randomIndex := rand.Intn(len(configs.Data))
	ctx.Set(SecurityID, configs.Data[randomIndex].ID)
	ctx.Set(SecurityName, configs.Data[randomIndex].Name)
	ctx.Set(SecurityValue, configs.Data[randomIndex].Value)
	ctx.Set(SecurityDescription, configs.Data[randomIndex].Description)

	return nil, nil
}
