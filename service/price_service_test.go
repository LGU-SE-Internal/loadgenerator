package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateQueryDeletePrice(t *testing.T) {
	cli, _ := GetAdminClient()

	MockedRouteId := faker.UUIDHyphenated()
	MockedTrainType := "HighSpeed"

	// Create a new price config
	createReq := PriceConfig{
		Id:        faker.UUIDHyphenated(),
		RouteId:   MockedRouteId,
		TrainType: MockedTrainType,
		Price:     100.0,
	}
	createResp, err := cli.CreateNewPriceConfig(createReq)
	if err != nil {
		t.Errorf("CreateNewPriceConfig failed: %v", err)
	}
	t.Logf("CreateNewPriceConfig response: %+v", createResp)

	// Update the price config
	updateReq := PriceConfig{
		Id:        createReq.Id,
		RouteId:   MockedRouteId,
		TrainType: MockedTrainType,
		Price:     150.0, // New price
	}
	updateResp, err := cli.UpdatePriceConfig(updateReq)
	if err != nil {
		t.Errorf("UpdatePriceConfig failed: %v", err)
	}
	t.Logf("UpdatePriceConfig response: %+v", updateResp)

	// Query price config by route ID and train type
	priceByRouteAndTrain, err := cli.FindByRouteIdAndTrainType(MockedRouteId, MockedTrainType)
	if err != nil {
		t.Errorf("FindByRouteIdAndTrainType failed: %v", err)
	}
	t.Logf("FindByRouteIdAndTrainType response: %+v", priceByRouteAndTrain)

	// Query all price configs
	allPriceConfigs, err := cli.FindAllPriceConfig()
	if err != nil {
		t.Errorf("FindAllPriceConfig failed: %v", err)
	}
	t.Logf("FindAllPriceConfig response: %+v", allPriceConfigs)

	// Delete the price config
	deleteResp, err := cli.DeletePriceConfig(createReq.Id)
	if err != nil {
		t.Errorf("DeletePriceConfig failed: %v", err)
	}
	t.Logf("DeletePriceConfig response: %+v", deleteResp)
}
