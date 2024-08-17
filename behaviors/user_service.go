package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

func QueryUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.UserService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	allUsersResp, err := cli.GetUserByUserId(ctx.Get(UserId).(string))
	if err != nil {
		log.Errorf("Request failed, err1 %s", err)
		return nil, err
	}
	if allUsersResp.Status != 1 {
		log.Errorf("Expected status 200, got %d", allUsersResp.Status)
		return nil, err
	}

	ctx.Set(UserName, allUsersResp.Data.UserName)
	ctx.Set(Password, allUsersResp.Data.Password)
	ctx.Set(Gender, allUsersResp.Data.Gender)
	ctx.Set(DocumentNum, allUsersResp.Data.DocumentNum)
	ctx.Set(DocumentType, allUsersResp.Data.DocumentType)
	ctx.Set(Email, allUsersResp.Data.Email)

	return nil, nil
}
