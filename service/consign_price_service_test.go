package service

import (
	"testing"
)

func TestSvcImpl_GetModifyDeletePriceConfig(t *testing.T) {
	cli, _ := GetAdminClient()
	//MockedID := faker.UUIDHyphenated()

	// Get price by weight and region
	priceByWeightAndRegion, err := cli.GetPriceByWeightAndRegion("50", "true")
	if err != nil {
		t.Errorf("GetPriceByWeightAndRegion failed: %v", err)
	}
	t.Logf("GetPriceByWeightAndRegion response: %+v", priceByWeightAndRegion)

	// Get price info
	priceInfo, err := cli.GetPriceInfo()
	if err != nil {
		t.Errorf("GetPriceInfo failed: %v", err)
	}
	t.Logf("GetPriceInfo response: %+v", priceInfo)

	// Get price config
	priceConfig, err := cli.GetPriceConfig()
	if err != nil {
		t.Errorf("GetPriceConfig failed: %v", err)
	}
	t.Logf("GetPriceConfig response: %+v", priceConfig)

	// Get the modified price config to verify changes
	getPriceConfig, err := cli.GetPriceConfig()
	if err != nil {
		t.Errorf("getPriceConfig after modification failed: %v", err)
	}
	t.Logf("GetPriceConfig after modification response: %+v", getPriceConfig)

	// Modify price config
	var getID string
	var getIndex int
	var getInitialWeight float64
	var getInitialPrice float64
	var getWithinPrice float64
	var getBeyondPrice float64
	if getPriceConfig.Data.Id != "" {
		getID = getPriceConfig.Data.Id
		getIndex = getPriceConfig.Data.Index
		getInitialPrice = getPriceConfig.Data.InitialPrice
		getInitialPrice = getPriceConfig.Data.InitialPrice
		getWithinPrice = getPriceConfig.Data.WithinPrice
		getBeyondPrice = getPriceConfig.Data.BeyondPrice
	}

	modifyResp, err := cli.ModifyPriceConfig(&ConsignPrice{
		ID:            getID,
		Index:         getIndex,
		InitialWeight: getInitialWeight,
		InitialPrice:  getInitialPrice,
		WithinPrice:   getWithinPrice,
		BeyondPrice:   getBeyondPrice,
	})
	if err != nil {
		t.Errorf("ModifyPriceConfig failed: %v", err)
	}
	t.Logf("ModifyPriceConfig response: %+v", modifyResp)
}
