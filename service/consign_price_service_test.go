package service

import (
	"math/rand"
	"testing"
)

func TestSvcImpl_GetModifyDeletePriceConfig(t *testing.T) {
	cli, _ := GetAdminClient()
	var consignSvc ConsignPriceService = cli

	// Mock Data
	MockedWeight := "8"
	MockedIsWithinWeight := "true"
	// Get price by weight and region
	priceByWeightAndRegion, err := consignSvc.GetPriceByWeightAndRegion(MockedWeight, MockedIsWithinWeight)
	if err != nil {
		t.Errorf("GetPriceByWeightAndRegion failed: %v", err)
	}
	if priceByWeightAndRegion.Status != 1 {
		t.Errorf("GetPriceByWeightAndRegion failed")
	}
	t.Logf("GetPriceByWeightAndRegion response: %+v", priceByWeightAndRegion)
	//existedConsignPrice := priceByWeightAndRegion.Data

	// Get price info
	priceInfo, err := cli.GetPriceInfo()
	if err != nil {
		t.Errorf("GetPriceInfo failed: %v", err)
	}
	if priceInfo.Status != 1 {
		t.Errorf("GetPriceInfo failed")
	}

	// Get price config
	priceConfig, err := cli.GetPriceConfig()
	if err != nil {
		t.Errorf("GetPriceConfig failed: %v", err)
	}
	if priceConfig.Status != 1 {
		t.Errorf("GetPriceConfig failed")
	}
	t.Logf("GetPriceConfig response: %+v", priceConfig)

	// Modify price config
	var getID string
	var getIndex int
	var getInitialWeight float64
	var getInitialPrice float64
	var getWithinPrice float64
	var getBeyondPrice float64

	getID = "39f89515-2d68-4ffb-9214-3c25a73da65f" // The ID here should be updated every redeploy the service since it is randomly generated when deploying.
	getIndex = 0
	getInitialWeight = rand.Float64()
	getInitialPrice = rand.Float64()
	getWithinPrice = rand.Float64()
	getBeyondPrice = rand.Float64()

	consignPrice := ConsignPrice{
		ID:            getID,
		Index:         getIndex,
		InitialWeight: getInitialWeight,
		InitialPrice:  getInitialPrice,
		WithinPrice:   getWithinPrice,
		BeyondPrice:   getBeyondPrice,
	}

	modifyResp, err := cli.ModifyPriceConfig(&consignPrice)
	if err != nil {
		t.Errorf("ModifyPriceConfig failed: %v", err)
	}
	if modifyResp.Status != 1 {
		t.Errorf("ModifyPriceConfig failed")
	}
	t.Logf("ModifyPriceConfig response: %+v", modifyResp)
}
