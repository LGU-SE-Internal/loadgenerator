package service

import (
	"math/rand"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateQueryDeletePrice(t *testing.T) {
	cli, _ := GetAdminClient()
	var priceSvc PriceService = cli

	MockedID := faker.UUIDHyphenated()

	var MockedTrainType string
	// 定义可能的开头字母
	letters := []rune{'Z', 'T', 'K', 'G', 'D'}
	// 随机选择一个字母
	startLetter := letters[rand.Intn(len(letters))]
	if startLetter == 'G' {
		if rand.Intn(2) == 0 {
			MockedTrainType = "GaoTieOne"
		} else {
			MockedTrainType = "GaoTieTwo"
		}

	} else if startLetter == 'Z' {
		MockedTrainType = "ZhiDa"
	} else if startLetter == 'T' {
		MockedTrainType = "TeKuai"
	} else if startLetter == 'K' {
		MockedTrainType = "KuaiSu"
	} else if startLetter == 'D' {
		MockedTrainType = "DongCheOne"
	}

	MockedRouteID := faker.UUIDHyphenated()

	MockedBasicPriceRate := rand.Float64()

	MockedFirstClassPriceRate := rand.Float64()

	// Create a new price config
	createReq := &PriceConfig{
		ID:                  MockedID,
		TrainType:           MockedTrainType,
		RouteID:             MockedRouteID,
		BasicPriceRate:      MockedBasicPriceRate,
		FirstClassPriceRate: MockedFirstClassPriceRate,
	}

	createResp, err := priceSvc.CreateNewPriceConfig(createReq)
	if err != nil {
		t.Errorf("CreateNewPriceConfig failed: %v", err)
	}
	if createResp.Msg == "Already exists" {
		t.Log("price found, skip")
		t.Skip()
	}
	if createResp.Data.Id != createReq.ID {
		t.Errorf("CreateNewPriceConfig ID failed: got %v, want %v", createResp.Data.Id, createReq.ID)
	}
	if createResp.Data.TrainType != createReq.TrainType {
		t.Errorf("CreateNewPriceConfig TrainType failed: got %v, want %v", createResp.Data.TrainType, createReq.TrainType)
	}
	if createResp.Data.RouteId != createReq.RouteID {
		t.Errorf("CreateNewPriceConfig RouteId failed: got %v, want %v", createResp.Data.RouteId, createReq.RouteID)
	}
	if createResp.Data.BasicPriceRate != createReq.BasicPriceRate {
		t.Errorf("CreateNewPriceConfig BasicPriceRate failed: got %v, want %v", createResp.Data.BasicPriceRate, createReq.BasicPriceRate)
	}
	if createResp.Data.FirstClassPriceRate != createReq.FirstClassPriceRate {
		t.Errorf("CreateNewPriceConfig FirstClassPriceRate failed: got %v, want %v", createResp.Data.FirstClassPriceRate, createReq.FirstClassPriceRate)
	}
	t.Logf("CreateNewPriceConfig response: %+v", createResp)
	existedPrice := createResp.Data

	// QueryTraintype all price configs
	allPriceConfigs, err1 := priceSvc.FindAllPriceConfig()
	if err1 != nil {
		t.Errorf("FindAllPriceConfig failed: %v", err1)
	}
	found := false
	for _, price := range allPriceConfigs.Data {
		if price.Id == existedPrice.Id {
			found = true
		}
	}
	if !found {
		t.Errorf("Request failed, station not found")
	}
	t.Logf("FindAllPriceConfig response: %+v", allPriceConfigs)

	// Update the price config
	updateReq := &PriceConfig{
		ID:                  MockedID,
		TrainType:           MockedTrainType,
		RouteID:             MockedRouteID,
		BasicPriceRate:      rand.Float64(),
		FirstClassPriceRate: rand.Float64(),
	}
	updateResp, err := priceSvc.UpdatePriceConfig(updateReq)
	if err != nil {
		t.Errorf("UpdatePriceConfig failed: %v", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("UpdatePriceConfig status failed: got %v, want %v", updateResp.Status, 1)
	}
	t.Logf("UpdatePriceConfig response: %+v", updateResp)

	// QueryTraintype price config by route ID and train type
	priceByRouteAndTrain, err := priceSvc.FindByRouteIdAndTrainType(existedPrice.RouteId, existedPrice.TrainType)
	if err != nil {
		t.Errorf("FindByRouteIdAndTrainType failed: %v", err)
	}
	if priceByRouteAndTrain.Status != 1 {
		t.Errorf("priceByRouteAndTrain.Status != 1")
	}
	if priceByRouteAndTrain.Data.Id != existedPrice.Id {
		t.Errorf("priceByRouteAndTrain.Data.Id != existedPrice.Id")
	}
	t.Logf("FindByRouteIdAndTrainType response: %+v", priceByRouteAndTrain)

	// Delete the price config
	deleteResp, err := priceSvc.DeletePriceConfig(createReq.ID)
	if err != nil {
		t.Errorf("DeletePriceConfig failed: %v", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("deleteResp.Status != 1")
	}
	t.Logf("DeletePriceConfig response: %+v", deleteResp)
}
