package service

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestUserService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	MockedID := faker.UUIDHyphenated()
	MockedUserName := faker.Username()

	// Test RegisterUser
	input := &UserDto{
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

	// Test GetAllUsers
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

	// Test UpdateUser
	updateInput := &UserDto{
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

	// Test DeleteUser
	deleteResp, err5 := cli.DeleteUser(MockedID)
	if err5 != nil {
		t.Errorf("Request failed, err5 %s", err5)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Expected status 200, got %d", deleteResp.Status)
	}
}
