package behaviors

import (
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"time"
)

type CreateUserBehavior struct{}

func (o *CreateUserBehavior) Run(cli *service.SvcImpl) {
	RegisterResp, err := cli.ReqUserCreate(&service.UserCreateInfoReq{
		Password: faker.Password(),
		UserName: faker.Username(),
		UserId:   uuid.New().String(),
	})
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)
	// login
	_, err = cli.ReqUserLogin(&service.UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	if err != nil {
		return
	}
	time.Sleep(2 * time.Second)
	// delete
	_, err = cli.ReqUserDelete(RegisterResp.Data.UserId)
	if err != nil {
		return
	}
}
