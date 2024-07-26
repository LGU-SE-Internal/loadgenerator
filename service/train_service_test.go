package service

import (
	"math/rand"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestTrainService_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	var trainSvc TrainService = cli

	// Mock data
	MockedID := faker.UUIDHyphenated()
	//options := []string{"GaoTieOne", "GaoTieTwo", "DongCheOne", "ZhiDa", "TeKuai", "KuaiSu", "QianNianSunHao"}
	//selectedName := RandomSelectString(options)
	//MockedName := selectedName
	//MockedName := faker.Name()
	MockedName := GenerateTrainTypeName()
	MockedEconomyClass := 2147483647 // MAX Value
	MockedConfortClass := 2147483647 // Max Value
	MockedAverageSpeed := 250 + rand.Intn(20)
	// input
	trainType := TrainType{
		AverageSpeed: MockedAverageSpeed,
		ConfortClass: MockedConfortClass,
		EconomyClass: MockedEconomyClass,
		Id:           MockedID,
		Name:         MockedName,
	}

	// Create Test
	createResp, err := trainSvc.Create(&trainType)
	if err != nil {
		t.Errorf("Create request failed, err %s", err)
	}
	if createResp.Status != 1 {
		t.Errorf("Create failed: %s", createResp.Msg)
	}
	//if createResp.Data.Id != trainType.Id {
	//	t.Errorf("Create failed: %s", createResp.Data.Id)
	//}
	if createResp.Data.Name != trainType.Name {
		t.Errorf("Create failed: %s", createResp.Data.Name)
	}
	if createResp.Data.EconomyClass != trainType.EconomyClass {
		t.Errorf("Create failed: %d", createResp.Data.EconomyClass)
	}
	if createResp.Data.ConfortClass != trainType.ConfortClass {
		t.Errorf("Create failed: %d", createResp.Data.ConfortClass)
	}
	if createResp.Data.AverageSpeed != trainType.AverageSpeed {
		t.Errorf("Create failed: %d", createResp.Data.AverageSpeed)
	}
	existedtrainType := trainType

	// Query Test
	resp, err := trainSvc.Query()
	if err != nil {
		t.Errorf("Request failed, err %s; while response: %v", err, resp)
	}
	//t.Logf("Query returned results: %v", resp)

	// Query all
	allTrainTypes, err := trainSvc.Query()
	if err != nil {
		t.Errorf("Query all request failed, err %s", err)
	}
	if allTrainTypes.Status != 1 {
		t.Errorf("allTrainTypes.Status != 1")
	}
	if len(allTrainTypes.Data) == 0 {
		t.Errorf("Query all returned no results")
	}
	found := false
	for _, trainTypeElement := range allTrainTypes.Data {
		if trainTypeElement.Id == createResp.Data.Id &&
			trainTypeElement.Name == existedtrainType.Name &&
			trainTypeElement.AverageSpeed == existedtrainType.AverageSpeed &&
			trainTypeElement.ConfortClass == existedtrainType.ConfortClass &&
			trainTypeElement.EconomyClass == existedtrainType.ConfortClass {
			found = true
		}
	}
	if !found {
		t.Errorf("Query all not get the corresponsing result, whcih means 'Creation Fails'")
	}

	// Test Update
	UpdatedAverageSpeed := 275 + rand.Intn(10)
	updateTrainType := TrainType{
		Id:           createResp.Data.Id,
		Name:         trainType.Name,
		EconomyClass: trainType.EconomyClass,
		ConfortClass: trainType.ConfortClass,
		AverageSpeed: UpdatedAverageSpeed,
	}
	updateResp, err := trainSvc.Update(&updateTrainType)
	if err != nil {
		t.Errorf("Update request failed, err %s", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("Update failed: %s", updateResp.Msg)
	}

	// Test Retrieve by Id
	retrieveResp, err := trainSvc.Retrieve(createResp.Data.Id)
	if err != nil {
		t.Errorf("Retrieve request failed, err %s", err)
	}
	//if len(retrieveResp.Data) == 0 {
	//	t.Errorf("Retrieve returned no result")
	//}
	if retrieveResp.Data == nil {
		t.Errorf("Retrieve returned no result")
	}

	// Test Retrieve by Name
	retrieveByNameResp, err := trainSvc.RetrieveByName(createResp.Data.Name)
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

	// Test Delete
	//var deleteID string
	//if len(allTrainTypes.Data) > 0 {
	//	deleteID = allTrainTypes.Data[len(allTrainTypes.Data)-1].Id
	//} else {
	//	t.Errorf("Query all returned empty data")
	//}
	deleteResp, err := trainSvc.Delete(createResp.Data.Id)
	if err != nil {
		t.Errorf("Delete request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Delete failed: %s", deleteResp.Msg)
	}

}
