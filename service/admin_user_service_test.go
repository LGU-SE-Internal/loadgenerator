package service

import (
	"github.com/go-faker/faker/v4"
	"testing"
)

func TestSvcImpl_AddUpdateDeleteUser(t *testing.T) {
	cli, _ := GetAdminClient()
	MockedID := faker.UUIDHyphenated()
	MockedUserName := faker.Username()

	// Add User
	addResp, err := cli.AddUser(&UserDto{
		UserID:       MockedID,
		UserName:     MockedUserName,
		Password:     "77777",
		Gender:       1,
		DocumentType: 1,
		DocumentNum:  "admin",
		Email:        faker.Email(),
	})
	if err != nil {
		t.Errorf("AddUser failed: %v", err)
	}
	t.Logf("AddUser response: %+v", addResp)

	// Update User
	updateResp, err := cli.UpdateUser(&UserDto{
		UserID:       MockedID,
		UserName:     MockedUserName,
		Password:     "88888",
		Gender:       1,
		DocumentType: 1,
		DocumentNum:  "admin",
		Email:        faker.Email(),
	})
	if err != nil {
		t.Errorf("UpdateUser failed: %v", err)
	}
	t.Logf("UpdateUser response: %+v", updateResp)

	// Delete User
	deleteResp, err := cli.DeleteUser(MockedID)
	if err != nil {
		t.Errorf("DeleteUser failed: %v", err)
	}
	t.Logf("DeleteUser response: %+v", deleteResp)

	// Get All Users
	users, err := cli.GetAllUsers()
	if err != nil {
		t.Errorf("GetAllUsers failed: %v", err)
	}
	t.Logf("GetAllUsers response: %+v", users)
}
