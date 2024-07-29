package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

const (
	UserName = "username"
	Password = "password"
	UserId   = "userid"
)

var LoginChain *Chain

func init() {
	LoginChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		return nil, nil
	}))
	LoginChain.AddNextChain(NewChain(NewFuncNode(LoginAdmin)), 0.2)
	LoginChain.AddNextChain(NewChain(NewFuncNode(CreateUser), NewFuncNode(LoginNormal)), 0.8)
}
func LoginAdmin(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	// login
	_, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	return nil, err
}

func LoginNormal(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	_, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
		Password:         ctx.Get(Password).(string),
		UserName:         ctx.Get(UserName).(string),
		VerificationCode: "123",
	})
	return nil, err
}

func CreateUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	RegisterResp, err := cli.ReqUserCreate(&service.UserCreateInfoReq{
		Password: faker.Password(),
		UserName: faker.Username(),
		UserId:   uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}
	ctx.Set(UserName, RegisterResp.Data.UserName)
	ctx.Set(Password, RegisterResp.Data.Password)
	ctx.Set(UserId, RegisterResp.Data.UserId)
	return nil, nil
}
