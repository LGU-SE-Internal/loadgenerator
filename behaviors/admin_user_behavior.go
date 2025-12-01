package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// Admin User Behavior Chain - 管理用户的增删改查操作
var AdminUserChain *Chain

func init() {
	AdminUserChain = NewChain(
		NewFuncNode(LoginAdmin, "LoginAdmin"),
		NewFuncNode(AdminQueryAllUsers, "AdminQueryAllUsers"),
	)

	// 添加后续的用户管理操作链
	AdminUserChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddUser, "AdminAddUser"),
		NewFuncNode(AdminUpdateUser, "AdminUpdateUser"),
	), 0.5)

	AdminUserChain.AddNextChain(NewChain(
		NewFuncNode(AdminDeleteUser, "AdminDeleteUser"),
	), 0.5)
}

// AdminQueryAllUsers 查询所有用户
func AdminQueryAllUsers(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.AdminGetAllUsers()
	if err != nil {
		log.Errorf("AdminQueryAllUsers failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminQueryAllUsers returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	if len(resp.Data) > 0 {
		randomIndex := rand.Intn(len(resp.Data))
		ctx.Set(UserId, resp.Data[randomIndex].UserID)
		ctx.Set(UserName, resp.Data[randomIndex].UserName)
		ctx.Set(Gender, resp.Data[randomIndex].Gender)
		ctx.Set(DocumentType, resp.Data[randomIndex].DocumentType)
		ctx.Set(DocumentNum, resp.Data[randomIndex].DocumentNum)
		ctx.Set(Email, resp.Data[randomIndex].Email)
	}

	log.Infof("AdminQueryAllUsers returned %d users", len(resp.Data))
	return nil, nil
}

// AdminAddUser 添加用户
func AdminAddUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	user := &service.AdminUserDto{
		UserID:       uuid.New().String(),
		UserName:     faker.Username(),
		Password:     faker.Password(),
		Gender:       rand.Intn(2),
		DocumentType: rand.Intn(2),
		DocumentNum:  faker.CCNumber(),
		Email:        faker.Email(),
	}

	resp, err := cli.AdminAddUser(user)
	if err != nil {
		log.Errorf("AdminAddUser failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddUser returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(UserId, resp.Data.UserId)
	ctx.Set(UserName, resp.Data.UserName)
	ctx.Set(Password, resp.Data.Password)
	ctx.Set(Email, resp.Data.Email)

	log.Infof("AdminAddUser success: userId=%s, userName=%s", resp.Data.UserId, resp.Data.UserName)
	return nil, nil
}

// AdminUpdateUser 更新用户
func AdminUpdateUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	userId, ok := ctx.Get(UserId).(string)
	if !ok || userId == "" {
		log.Warn("No user ID found in context, skipping update")
		return nil, nil
	}

	userName, _ := ctx.Get(UserName).(string)
	if userName == "" {
		userName = faker.Username()
	}

	user := &service.AdminUserDto{
		UserID:       userId,
		UserName:     userName,
		Password:     faker.Password(),
		Gender:       rand.Intn(2),
		DocumentType: rand.Intn(2),
		DocumentNum:  faker.CCNumber(),
		Email:        faker.Email(),
	}

	resp, err := cli.AdminUpdateUser(user)
	if err != nil {
		log.Errorf("AdminUpdateUser failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminUpdateUser returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminUpdateUser success: userId=%s", userId)
	return nil, nil
}

// AdminDeleteUser 删除用户
func AdminDeleteUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	userId, ok := ctx.Get(UserId).(string)
	if !ok || userId == "" {
		log.Warn("No user ID found in context, skipping delete")
		return nil, nil
	}

	resp, err := cli.AdminDeleteUser(userId)
	if err != nil {
		log.Errorf("AdminDeleteUser failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminDeleteUser returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminDeleteUser success: userId=%s", userId)
	return nil, nil
}
