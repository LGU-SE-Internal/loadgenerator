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
	addResp, err := cli.AdminAddUser(&AdminUserDto{
		UserID:       MockedID,
		UserName:     MockedUserName,
		Password:     "77777",
		Gender:       1,
		DocumentType: 1,
		DocumentNum:  "admin",
		Email:        faker.Email(),
	})
	if err != nil {
		t.Errorf("AdminAddUser failed: %v", err)
	}
	t.Logf("AdminAddUser response: %+v", addResp)

	// Update User
	updateResp, err := cli.AdminUpdateUser(&AdminUserDto{
		UserID:       MockedID,
		UserName:     MockedUserName,
		Password:     "88888",
		Gender:       1,
		DocumentType: 1,
		DocumentNum:  "admin",
		Email:        faker.Email(),
	})
	if err != nil {
		t.Errorf("AdminUpdateUser failed: %v", err)
	}
	t.Logf("AdminUpdateUser response: %+v", updateResp)

	// Delete User
	deleteResp, err := cli.AdminDeleteUser(MockedID)
	if err != nil {
		t.Errorf("AdminDeleteUser failed: %v", err)
	}
	t.Logf("AdminDeleteUser response: %+v", deleteResp)

	// Get All Users
	users, err := cli.AdminGetAllUsers()
	if err != nil {
		t.Errorf("AdminGetAllUsers failed: %v", err)
	}
	t.Logf("AdminGetAllUsers response: %+v", users)
}
