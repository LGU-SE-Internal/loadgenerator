package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
)

func QueryConsignPric(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part

	return nil, nil
}

func CreateConsignPrice(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")

	}

	// TODO part

	return nil, nil
}
