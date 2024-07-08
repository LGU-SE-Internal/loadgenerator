package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestTrainService_FullIntegration(t *testing.T) {
	cli, _ := GetBasicClient()

	// Query Test
	resp, err := cli.Query()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	t.Logf("Query returned results: %v", resp)

	// Mock data
	MockedID := faker.UUIDHyphenated()
	MockedName := faker.Name()
	trainType := &TrainType{
		ID:           MockedID,
		Name:         MockedName,
		EconomyClass: 1,
		ConfortClass: 0,
		AverageSpeed: 200,
	}

	// Create Test
	createResp, err := cli.Create(trainType)
	if err != nil {
		t.Errorf("Create request failed, err %s", err)
	}
	if createResp.Status != 1 {
		t.Errorf("Create failed: %s", createResp.Msg)
	}

	// Query all
	allTrainTypes, err := cli.Query()
	if err != nil {
		t.Errorf("Query all request failed, err %s", err)
	}
	if len(allTrainTypes.Data) == 0 {
		t.Errorf("Query all returned no results")
	}

	var getId string
	var getName string
	if len(allTrainTypes.Data) > 0 {
		getId = allTrainTypes.Data[0].Id
		getName = allTrainTypes.Data[0].Name
	}

	// Test Update
	updateTrainType := &TrainType{
		ID:           getId,
		Name:         getName,
		EconomyClass: 7,
		ConfortClass: 8,
		AverageSpeed: 234,
	}
	updateResp, err := cli.Update(updateTrainType)
	if err != nil {
		t.Errorf("Update request failed, err %s", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("Update failed: %s", updateResp.Msg)
	}

	// Test Delete
	var deleteID string
	if len(allTrainTypes.Data) > 0 {
		deleteID = allTrainTypes.Data[len(allTrainTypes.Data)-1].Id
	} else {
		t.Errorf("Query all returned empty data")
	}

	deleteResp, err := cli.Delete(deleteID)
	if err != nil {
		t.Errorf("Delete request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Delete failed: %s", deleteResp.Msg)
	}

	// Test Retrieve by ID
	retrieveResp, err := cli.Retrieve(getId)
	if err != nil {
		t.Errorf("Retrieve request failed, err %s", err)
	}
	if retrieveResp == nil {
		t.Errorf("Retrieve returned no result")
	}

	// Test Retrieve by Name
	retrieveByNameResp, err := cli.RetrieveByName(getName)
	if err != nil {
		t.Errorf("Retrieve by name request failed, err %s", err)
	}
	if retrieveByNameResp == nil {
		t.Errorf("Retrieve by name returned no result")
	}

	// Test Retrieve by Names
	names := []string{"GaoTieOne", "GaoTieTwo", "DongCheOne"}
	retrieveByNamesResp, err := cli.RetrieveByNames(names)
	if err != nil {
		t.Errorf("Retrieve by names request failed, err %s", err)
	}
	if len(retrieveByNamesResp.Data) == 0 {
		t.Errorf("Retrieve by names returned no results")
	}
}
