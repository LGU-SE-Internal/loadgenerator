package service

import (
	"testing"
)

func TestSvcImpl_AddUpdateDeleteUser(t *testing.T) {
	cli, _ := GetBasicClient()

	// Add User
	addResp, err := cli.AddUser(&UserDto{
		//UserID       string `json:"userId"`
		//UserName     string `json:"userName"`
		//Password     string `json:"password"`
		//Gender       int    `json:"gender"`
		//DocumentType int    `json:"documentType"`
		//DocumentNum  string `json:"documentNum"`
		//Email        string `json:"email"`
	})
	if err != nil {
		t.Errorf("AddUser failed: %v", err)
	}
	t.Logf("AddUser response: %+v", addResp)

	// Update User
	updateResp, err := cli.UpdateUser(&UserDto{
		//UserID       string `json:"userId"`
		//UserName     string `json:"userName"`
		//Password     string `json:"password"`
		//Gender       int    `json:"gender"`
		//DocumentType int    `json:"documentType"`
		//DocumentNum  string `json:"documentNum"`
		//Email        string `json:"email"`
	})
	if err != nil {
		t.Errorf("UpdateUser failed: %v", err)
	}
	t.Logf("UpdateUser response: %+v", updateResp)

	// Delete User
	deleteResp, err := cli.DeleteUser("12345")
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
