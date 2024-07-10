package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestConfigService_FullIntegration(t *testing.T) {
	cli, _ := GetBasicClient() // Assuming GetBasicClient is implemented elsewhere

	// Query All Configs Test
	queryAllResp, err := cli.QueryAllConfigs()
	if err != nil {
		t.Errorf("QueryAllConfigs request failed, err %s", err)
	}
	t.Logf("QueryAllConfigs returned results: %v", queryAllResp)

	// Mock data
	mockedName := faker.Name()
	mockedValue := faker.Word()
	mockedDescription := faker.Sentence()
	config := &Config_config{
		Name:        mockedName,
		Value:       mockedValue,
		Description: mockedDescription,
	}

	// Create Config Test
	createResp, err := cli.CreateConfig(config)
	if err != nil {
		t.Errorf("CreateConfig request failed, err %s", err)
	}
	t.Logf("CreateConfig returned results: %v", createResp)

	// Query All Configs again to get the ID of the created config
	allConfigsResp, err := cli.QueryAllConfigs()
	if err != nil {
		t.Errorf("QueryAllConfigs request failed, err %s", err)
	}
	t.Logf("QueryAllConfigs returned results: %v", allConfigsResp)

	var GetName string
	var GetName2 string
	var GetValue string
	var GetDescription string
	if len(allConfigsResp.Data) > 0 {
		GetName = allConfigsResp.Data[0].Name
		GetName2 = allConfigsResp.Data[len(allConfigsResp.Data)-1].Name
		GetValue = allConfigsResp.Data[0].Value
		GetDescription = allConfigsResp.Data[0].Description
	}

	// Test Update Config
	updatedConfig := Config_config{
		Name:        GetName,
		Value:       GetValue,
		Description: GetDescription,
	}
	updateResp, err := cli.UpdateConfig(updatedConfig)
	if err != nil {
		t.Errorf("UpdateConfig request failed, err %s", err)
	}
	t.Logf("UpdateConfig returned results: %v", updateResp)

	// Test Retrieve Config by Name
	retrieveResp, err := cli.RetrieveConfig(GetName)
	if err != nil {
		t.Errorf("RetrieveConfig request failed, err %s", err)
	}
	t.Logf("RetrieveConfig returned results: %v", retrieveResp)

	// Test Delete Config
	deleteResp, err := cli.DeleteConfig_config_service(GetName2)
	if err != nil {
		t.Errorf("DeleteConfig request failed, err %s", err)
	}
	t.Logf("DeleteConfig returned results: %v", deleteResp)

}

func TestConfigService_FullIntegration_v2(t *testing.T) {
	cli, _ := GetBasicClient() // Assuming GetBasicClient is implemented elsewhere

	// Query All Configs Test
	queryAllResp, err := cli.QueryAllConfigs()
	if len(queryAllResp.Data) < 1 {
		t.Errorf("No configs found.")
	}
	if err != nil {
		t.Errorf("QueryAllConfigs request failed, err %s", err)
	}
	t.Logf("QueryAllConfigs returned results: %v", queryAllResp)

	// Mock data
	mockedName := faker.Name()
	mockedValue := faker.Word()
	mockedDescription := faker.Sentence()
	config := &Config_config{
		Name:        mockedName,
		Value:       mockedValue,
		Description: mockedDescription,
	}

	// Create Config Test
	createResp, err := cli.CreateConfig(config)
	if err != nil {
		t.Errorf("CreateConfig request failed, err %s", err)
	}
	t.Logf("CreateConfig returned results: %v", createResp)
	if mockedName != createResp.Data.Name {
		t.Errorf("CREATE: wrong config name. got %s, want %s", createResp.Data.Name, mockedName)
	}
	if mockedValue != createResp.Data.Value {
		t.Errorf("CREATE: wrong config value. got %s, want %s", createResp.Data.Value, mockedValue)
	}
	if mockedDescription != createResp.Data.Description {
		t.Errorf("CREATE: wrong config description. got %s, want %s", createResp.Data.Description, mockedDescription)
	}

	retrieveResp, err := cli.RetrieveConfig(createResp.Data.Name)
	if err != nil {
		t.Errorf("RetrieveConfig request failed, err %s", err)
	}
	if mockedName != retrieveResp.Data.Name {
		t.Errorf("GET: wrong config name. got %s, want %s", retrieveResp.Data.Name, mockedName)
	}
	if mockedValue != retrieveResp.Data.Value {
		t.Errorf("GET: wrong config value. got %s, want %s", retrieveResp.Data.Value, mockedValue)
	}
	if mockedDescription != retrieveResp.Data.Description {
		t.Errorf("GET: wrong config description. got %s, want %s", retrieveResp.Data.Description, mockedDescription)
	}

	// Update value
	createResp.Data.Value = faker.Word()
	// Test Update Config
	updatedConfig := createResp.Data
	updateResp, err := cli.UpdateConfig(updatedConfig)
	if err != nil {
		t.Errorf("UpdateConfig request failed, err %s", err)
	}
	t.Logf("UpdateConfig returned results: %v", updateResp)

	if createResp.Data.Name != updateResp.Data.Name {
		t.Errorf("GET: wrong config name. got %s, want %s", updateResp.Data.Name, createResp.Data.Name)
	}
	if createResp.Data.Value != updateResp.Data.Value {
		t.Errorf("GET: wrong config value. got %s, want %s", updateResp.Data.Value, createResp.Data.Value)
	}
	if createResp.Data.Description != updateResp.Data.Description {
		t.Errorf("GET: wrong config description. got %s, want %s", updateResp.Data.Description, createResp.Data.Description)
	}

	// Test Retrieve Config by Name
	retrieveResp, err = cli.RetrieveConfig(createResp.Data.Name)
	if err != nil {
		t.Errorf("RetrieveConfig request failed, err %s", err)
	}
	t.Logf("RetrieveConfig returned results: %v", retrieveResp)

	// Test Delete Config
	deleteResp, err := cli.DeleteConfig_config_service(createResp.Data.Name)
	if err != nil {
		t.Errorf("DeleteConfig request failed, err %s", err)
	}
	t.Logf("DeleteConfig returned results: %v", deleteResp)

}
