package service

import (
	"math/rand"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateQueryDeletePrice(t *testing.T) {
	cli, _ := GetAdminClient()

	MockedRouteId := faker.UUIDHyphenated()
	MockedTrainType := "HighSpeed"
	//MockedTrainId := faker.UUIDHyphenated()

	// Create a new price config
	createReq := &PriceConfig{
		Id:        MockedRouteId,
		TrainType: MockedTrainType,
		Price:     rand.Float64(),
	}
	createResp, err := cli.CreateNewPriceConfig(createReq)
	if err != nil {
		t.Errorf("CreateNewPriceConfig failed: %v", err)
	}
	t.Logf("CreateNewPriceConfig response: %+v", createResp)

	// Update the price config
	updateReq := &PriceConfig{
		Id:        MockedRouteId,
		TrainType: MockedTrainType,
		Price:     rand.Float64(),
	}
	updateResp, err := cli.UpdatePriceConfig(updateReq)
	if err != nil {
		t.Errorf("UpdatePriceConfig failed: %v", err)
	}
	t.Logf("UpdatePriceConfig response: %+v", updateResp)

	// Query all price configs
	allPriceConfigs, err := cli.FindAllPriceConfig()
	if err != nil {
		t.Errorf("FindAllPriceConfig failed: %v", err)
	}
	t.Logf("FindAllPriceConfig response: %+v", allPriceConfigs)

	// Query price config by route ID and train type
	priceByRouteAndTrain, err := cli.FindByRouteIdAndTrainType(MockedRouteId, MockedTrainType)
	if err != nil {
		t.Errorf("FindByRouteIdAndTrainType failed: %v", err)
	}
	t.Logf("FindByRouteIdAndTrainType response: %+v", priceByRouteAndTrain)

	// Delete the price config
	deleteResp, err := cli.DeletePriceConfig(createReq.Id)
	if err != nil {
		t.Errorf("DeletePriceConfig failed: %v", err)
	}
	t.Logf("DeletePriceConfig response: %+v", deleteResp)
}
