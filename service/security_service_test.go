package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateDeleteSecurityConfig(t *testing.T) {
	cli, _ := GetAdminClient()
	var securitySvc SecurityService = cli

	//MockedID := faker.UUIDHyphenated()
	MockedName := faker.Name()
	MockedValue := generateRandomNumberString()
	MockedDescription := generateDescription()

	// Mock input
	input := &SecurityConfig{
		//ID:          MockedID,
		Name:        MockedName,
		Value:       MockedValue,
		Description: MockedDescription,
	}

	// Add Security AdminConfig
	addResp, err := securitySvc.AddNewSecurityConfig(input)
	if err != nil {
		t.Errorf("AddNewSecurityConfig failed: %v", err)
	}
	if addResp.Status != 1 {
		t.Errorf("[Security Service]addResp.Status != 1")
	}
	if addResp.Msg == "Already exists" {
		t.Logf("[Security Service]addResp.Msg => Already exists")
		t.Skip()
	}
	isMatch := false
	if /*addResp.Data.ID == input.ID &&*/
	addResp.Data.Value == input.Value &&
		addResp.Data.Description == input.Description &&
		addResp.Data.Name == input.Name {
		isMatch = true
	}
	if !isMatch {
		t.Errorf("[Security Service]Except: %v, get %v", input, addResp.Data)
	}
	t.Logf("AddNewSecurityConfig response: %+v", addResp)
	existedSecurity := addResp.Data

	// Get All Security Configs
	configs, err3 := securitySvc.FindAllSecurityConfig()
	if err3 != nil {
		t.Errorf("FindAllSecurityConfig failed: %v", err3)
	}
	if configs.Status != 1 {
		t.Errorf("[Security Service]Status != 1")
	}
	found := false
	for _, security := range configs.Data {
		if security.ID == existedSecurity.ID &&
			security.Name == existedSecurity.Name &&
			security.Value == existedSecurity.Value &&
			security.Description == existedSecurity.Description {
			found = true
		}
	}
	if !found {
		t.Errorf("[Security Service]Cannot find existed security config")
	}
	t.Logf("FindAllSecurityConfig response: %+v", configs)

	// Update Security AdminConfig
	updateResp, err1 := securitySvc.ModifySecurityConfig(&SecurityConfig{
		ID:          existedSecurity.ID,
		Name:        existedSecurity.Name,
		Value:       generateRandomNumberString(),
		Description: generateDescription(),
	})
	if err1 != nil {
		t.Errorf("ModifySecurityConfig failed: %v", err1)
	}
	if updateResp.Status != 1 {
		t.Errorf("[Security Service]Status != 1")
	}
	t.Logf("ModifySecurityConfig response: %+v", updateResp)

	// Delete Security AdminConfig
	deleteResp, err2 := cli.DeleteSecurityConfig(existedSecurity.ID)
	if err2 != nil {
		t.Errorf("DeleteSecurityConfig failed: %v", err2)
	}
	if deleteResp.Status != 1 {
		t.Errorf("[Security Service]Status != 0")
	}
	t.Logf("DeleteSecurityConfig response: %+v", deleteResp)

}
