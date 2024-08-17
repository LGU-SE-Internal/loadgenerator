package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

// FoodBehaviorChain
func QueryFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype all
	allFoodOrders, err := cli.FindAllFoodOrder()
	if err != nil {
		log.Errorf("FindAllFoodOrder request failed, err %s", err)
		return nil, err
	}
	if len(allFoodOrders.Data) == 0 {
		log.Errorf("FindAllFoodOrder returned empty results")
		return nil, err
	}
	if allFoodOrders.Status != 1 {
		log.Errorf("FindAllFoodOrder failed: %v", allFoodOrders.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(allFoodOrders.Data))
	//ctx.Set(OrderId, allFoodOrders.Data[randomIndex].OrderId)
	ctx.Set(FoodType, allFoodOrders.Data[randomIndex].FoodType)
	//ctx.Set(StationName, allFoodOrders.Data[randomIndex].StationName)
	ctx.Set(StoreName, allFoodOrders.Data[randomIndex].StoreName)
	ctx.Set(FoodName, allFoodOrders.Data[randomIndex].FoodName)
	ctx.Set(Price, allFoodOrders.Data[randomIndex].Price)

	return nil, nil
}

func CreateFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedOrderID := ctx.Get(OrderId).(string)
	MockedID := faker.UUIDHyphenated()
	foodOrder := service.FoodOrder{
		ID:          MockedID,
		OrderID:     MockedOrderID,
		FoodType:    rand.Intn(1),
		FoodName:    generateRandomFood(),
		StationName: ctx.Get(StationName).(string),
		StoreName:   ctx.Get(StoreName).(string),
		Price:       ctx.Get(Price).(float64),
	}

	// Create Test
	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		log.Errorf("NewCreateFoodOrder request failed, err %s", err)
		return nil, err
	}
	if newCreateResp.Status != 1 {
		log.Errorf("NEwCreateFoodOrder failed")
		return nil, err
	}

	ctx.Set(OrderId, newCreateResp.Data.OrderId)
	ctx.Set(FoodType, newCreateResp.Data.FoodType)
	ctx.Set(StoreName, newCreateResp.Data.StoreName)
	ctx.Set(FoodName, newCreateResp.Data.FoodName)
	ctx.Set(Price, newCreateResp.Data.Price)

	return nil, nil
}
