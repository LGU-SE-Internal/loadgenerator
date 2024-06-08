package service

import (
	"github.com/google/uuid"
	"testing"
)

func GetBasicClient() (*SvcImpl, string) {
	cli := NewSvcClients()
	loginResp, _ := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "111111",
		UserName:         "fdse_microservice",
		VerificationCode: "123",
	})
	return cli, loginResp.Data.UserId
}
func GetAdminClient() (*SvcImpl, string) {
	cli := NewSvcClients()
	loginResp, _ := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	return cli, loginResp.Data.UserId
}

func TestSvcImpl_ReqUserCreate(t *testing.T) {
	// create
	cli := NewSvcClients()
	RegisterResp, err := cli.ReqUserCreate(&UserCreateInfoReq{
		Password: "testpasswd",
		UserName: "testuser",
		UserId:   uuid.New().String(),
	})
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if RegisterResp.Status != 1 {
		t.Errorf("RegisterResp.Status != 1")
	}

	// login
	loginResp, err := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	if err != nil {
		t.Error(err)
	}
	if loginResp.Status != 1 {
		t.Errorf("RegisterResp.Status != 1")
	}
	if loginResp.Data.Username != "admin" {
		t.Errorf("RegisterResp.Data.Username != \"admin\"")
	}

	// delete
	deleteResp, err := cli.ReqUserDelete(RegisterResp.Data.UserId)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("deleteResp.Status != 1")
	}
}
