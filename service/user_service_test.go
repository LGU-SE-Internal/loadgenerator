package service

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUserService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	MockedID := faker.UUIDHyphenated()
	MockedUserName := faker.Username()

	// Test RegisterUser
	input := &AdminUserDto{
		UserID:       MockedID,
		UserName:     MockedUserName,
		Password:     "123456",
		Gender:       1,
		DocumentType: 1,
		DocumentNum:  "basic",
		Email:        faker.Email(),
	}

	resp, err := cli.RegisterUser(input)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("Expected status 201, got %d", resp.Status)
	}

	// Test AdminGetAllUsers
	allUsersResp, err1 := cli.GetAllUsers()
	if err1 != nil {
		t.Errorf("Request failed, err1 %s", err1)
	}
	if allUsersResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", allUsersResp.Status)
	}

	// Test GetUserByUserName
	userByNameResp, err2 := cli.GetUserByUserName(MockedUserName)
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}
	if userByNameResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", userByNameResp.Status)
	}

	// Test GetUserByUserId
	userByIdResp, err3 := cli.GetUserByUserId(MockedID)
	if err3 != nil {
		t.Errorf("Request failed, err3 %s", err3)
	}
	if userByIdResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", userByIdResp.Status)
	}

	// Test AdminUpdateUser
	updateInput := &AdminUserDto{
		UserID:       MockedID,
		UserName:     MockedUserName,
		Password:     "654321",
		Gender:       1,
		DocumentType: 1,
		DocumentNum:  "basic",
		Email:        faker.Email(),
	}
	updateResp, err4 := cli.UpdateUser(updateInput)
	if err4 != nil {
		t.Errorf("Request failed, err4 %s", err4)
	}
	if updateResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", updateResp.Status)
	}

	// Test AdminDeleteUser
	deleteResp, err5 := cli.DeleteUser(MockedID)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", deleteResp.Status)
	}
}

func TestUserService_FullIntegration_v2(t *testing.T) {
	cli, _ := GetAdminClient()
	var userSvc UserService = cli

	// 使用当前时间作为随机数生成器的种子，确保每次运行程序时得到不同的结果
	rand.Seed(time.Now().UnixNano())

	// Test RegisterUser
	input := &AdminUserDto{
		UserID:       uuid.NewString(),
		UserName:     faker.Name(),
		Password:     faker.Password(),
		Gender:       rand.Intn(2),
		DocumentType: rand.Intn(2),
		DocumentNum:  strconv.Itoa(rand.Intn(9999)),
		Email:        faker.Email(),
	}

	resp, err := cli.RegisterUser(input)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if resp.Status != 1 {
		t.Errorf("Expected status 201, got %d", resp.Status)
	}
	returnedUser := resp.Data
	if returnedUser.UserID != input.UserID {
		t.Errorf("Expected user id %s, got %s", input.UserID, returnedUser.UserID)
		t.Skip()
	}
	if returnedUser.UserName != input.UserName {
		t.Errorf("Expected user name %s, got %s", input.UserName, returnedUser.UserName)
		t.Skip()
	}
	if returnedUser.Password != input.Password {
		t.Errorf("Expected password %s, got %s", input.Password, returnedUser.Password)
		t.Skip()
	}
	if returnedUser.Gender != input.Gender {
		t.Errorf("Expected gender %d, got %d", input.Gender, returnedUser.Gender)
		t.Skip()
	}
	if returnedUser.DocumentType != input.DocumentType {
		t.Errorf("Expected document type %d, got %d", input.DocumentType, returnedUser.DocumentType)
		t.Skip()
	}
	if returnedUser.DocumentNum != input.DocumentNum {
		t.Errorf("Expected document num %s, got %s", input.DocumentNum, returnedUser.DocumentNum)
		t.Skip()
	}
	if returnedUser.Email != input.Email {
		t.Errorf("Expected email %s, got %s", input.Email, returnedUser.Email)
		t.Skip()
	}

	// Test AdminGetAllUsers
	allUsersResp, err1 := userSvc.GetAllUsers()
	if err1 != nil {
		t.Errorf("Request failed, err1 %s", err1)
	}
	if allUsersResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", allUsersResp.Status)
	}
	if len(allUsersResp.Data) == 0 {
		t.Errorf("Response should not be empty")
	}

	gbiUserResp, err2 := userSvc.GetUserByUserId(returnedUser.UserID)
	if err2 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}

	returnedUser = gbiUserResp.Data
	if returnedUser.UserID != input.UserID {
		t.Errorf("Expected user id %s, got %s", input.UserID, returnedUser.UserID)
		t.Skip()
	}
	if returnedUser.UserName != input.UserName {
		t.Errorf("Expected user name %s, got %s", input.UserName, returnedUser.UserName)
		t.Skip()
	}
	if returnedUser.Password != input.Password {
		t.Errorf("Expected password %s, got %s", input.Password, returnedUser.Password)
		t.Skip()
	}
	if returnedUser.Gender != input.Gender {
		t.Errorf("Expected gender %d, got %d", input.Gender, returnedUser.Gender)
		t.Skip()
	}
	if returnedUser.DocumentType != input.DocumentType {
		t.Errorf("Expected document type %d, got %d", input.DocumentType, returnedUser.DocumentType)
		t.Skip()
	}
	if returnedUser.DocumentNum != input.DocumentNum {
		t.Errorf("Expected document num %s, got %s", input.DocumentNum, returnedUser.DocumentNum)
		t.Skip()
	}
	if returnedUser.Email != input.Email {
		t.Errorf("Expected email %s, got %s", input.Email, returnedUser.Email)
		t.Skip()
	}

	gbnUserResp, err3 := userSvc.GetUserByUserName(returnedUser.UserName)
	if err3 != nil {
		t.Errorf("Request failed, err2 %s", err2)
	}

	returnedUser = gbnUserResp.Data
	if returnedUser.UserID != input.UserID {
		t.Errorf("Expected user id %s, got %s", input.UserID, returnedUser.UserID)
		t.Skip()
	}
	if returnedUser.UserName != input.UserName {
		t.Errorf("Expected user name %s, got %s", input.UserName, returnedUser.UserName)
		t.Skip()
	}
	if returnedUser.Password != input.Password {
		t.Errorf("Expected password %s, got %s", input.Password, returnedUser.Password)
		t.Skip()
	}
	if returnedUser.Gender != input.Gender {
		t.Errorf("Expected gender %d, got %d", input.Gender, returnedUser.Gender)
		t.Skip()
	}
	if returnedUser.DocumentType != input.DocumentType {
		t.Errorf("Expected document type %d, got %d", input.DocumentType, returnedUser.DocumentType)
		t.Skip()
	}
	if returnedUser.DocumentNum != input.DocumentNum {
		t.Errorf("Expected document num %s, got %s", input.DocumentNum, returnedUser.DocumentNum)
		t.Skip()
	}
	if returnedUser.Email != input.Email {
		t.Errorf("Expected email %s, got %s", input.Email, returnedUser.Email)
		t.Skip()
	}

	input = &AdminUserDto{
		UserID:       input.UserID,
		UserName:     faker.Name(),
		Password:     faker.Password(),
		Gender:       rand.Intn(2),
		DocumentType: rand.Intn(2),
		DocumentNum:  strconv.Itoa(rand.Intn(9999)),
		Email:        faker.Email(),
	}

	updateResp, err4 := userSvc.UpdateUser(input)
	if err4 != nil {
		t.Errorf("Request failed, err4 %s", err4)
	}
	returnedUser = updateResp.Data

	if returnedUser.UserID != input.UserID {
		t.Errorf("Expected user id %s, got %s", input.UserID, returnedUser.UserID)
		t.Skip()
	}
	if returnedUser.UserName != input.UserName {
		t.Errorf("Expected user name %s, got %s", input.UserName, returnedUser.UserName)
		t.Skip()
	}
	if returnedUser.Password != input.Password {
		t.Errorf("Expected password %s, got %s", input.Password, returnedUser.Password)
		t.Skip()
	}
	if returnedUser.Gender != input.Gender {
		t.Errorf("Expected gender %d, got %d", input.Gender, returnedUser.Gender)
		t.Skip()
	}
	if returnedUser.DocumentType != input.DocumentType {
		t.Errorf("Expected document type %d, got %d", input.DocumentType, returnedUser.DocumentType)
		t.Skip()
	}
	if returnedUser.DocumentNum != input.DocumentNum {
		t.Errorf("Expected document num %s, got %s", input.DocumentNum, returnedUser.DocumentNum)
		t.Skip()
	}
	if returnedUser.Email != input.Email {
		t.Errorf("Expected email %s, got %s", input.Email, returnedUser.Email)
		t.Skip()
	}

	deleteResp, err5 := userSvc.DeleteUser(input.UserID)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Expected status 1, got %d", deleteResp.Status)
	}
	if deleteResp.Msg != "DELETE SUCCESS" {
		t.Errorf("Expected msg %s, got %s", "DELETE SUCCESS", deleteResp.Msg)
	}

}
