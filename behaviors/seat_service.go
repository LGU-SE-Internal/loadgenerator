package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
)

func QuerySeatInfo(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.SeatService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	_ = cli

	ctx.Set(SeatClass, 2)
	return nil, nil
}
