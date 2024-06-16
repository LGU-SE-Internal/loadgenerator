package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateDeleteSecurityConfig(t *testing.T) {
	cli, _ := GetAdminClient()
	MockedID := faker.UUIDHyphenated()
	MockedName := faker.Name()

	// Mock input
	input := &SecurityConfig{
		ID:          MockedID,
		Name:        MockedName,
		Value:       "2147483647",
		Description: "Made in haven",
	}

	// Add Security Config
	addResp, err := cli.AddNewSecurityConfig(input)
	if err != nil {
		t.Errorf("AddNewSecurityConfig failed: %v", err)
	}
	t.Logf("AddNewSecurityConfig response: %+v", addResp)

	// Update Security Config
	updateResp, err1 := cli.ModifySecurityConfig(&SecurityConfig{
		ID:          "cf3cf953-7f67-4936-b8bc-e7e1720f26e2",
		Name:        "Prof. Ardella Schinner",
		Value:       "7777777777",
		Description: "Made in haven",
	})
	if err1 != nil {
		t.Errorf("ModifySecurityConfig failed: %v", err1)
	}
	t.Logf("ModifySecurityConfig response: %+v", updateResp)

	// Get All Security Configs
	configs, err3 := cli.FindAllSecurityConfig()
	if err3 != nil {
		t.Errorf("FindAllSecurityConfig failed: %v", err3)
	}
	t.Logf("FindAllSecurityConfig response: %+v", configs)

	// Delete Security Config
	var deleteID string
	if len(configs.Data) > 0 {
		deleteID = configs.Data[len(configs.Data)-1].ID
	} else {
		t.Errorf("allContacts.Data is empty")
	}
	deleteResp, err2 := cli.DeleteSecurityConfig(deleteID)
	if err2 != nil {
		t.Errorf("DeleteSecurityConfig failed: %v", err2)
	}
	t.Logf("DeleteSecurityConfig response: %+v", deleteResp)

}
