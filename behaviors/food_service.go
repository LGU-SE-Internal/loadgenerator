package behaviors

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

// FoodBehaviorChain
func QueryFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.FoodService)
	if !ok {
		return nil, fmt.Errorf("FoodService client not found in context")
	}

	TheDate := time.Now().Format("2006-01-02")
	TheStartStation := ctx.Get(StartStation).(string)
	TheEndStation := ctx.Get(EndStation).(string)
	TheTripID := ctx.Get(TripID).(string)

	var foodSvc service.FoodService = cli
	allFood, err := foodSvc.GetAllFood(TheDate, TheStartStation, TheEndStation, TheTripID)
	if err != nil {
		log.Errorf("Failed to retrieve all food options: %v | Parameters: date=%s, startStation=%s, endStation=%s, tripID=%s", err, TheDate, TheStartStation, TheEndStation, TheTripID)
		return &NodeResult{false}, err
	}

	if allFood.Status != 1 {
		log.Warnf("Food service returned non-success status: %d | Response: %+v | Parameters: date=%s, startStation=%s, endStation=%s, tripID=%s", allFood.Status, allFood, TheDate, TheStartStation, TheEndStation, TheTripID)
		return &NodeResult{false}, errors.New("food service responded with status error")
	}
	foodType := rand.Int()%2 + 1 // 1: train food, 2: food store
	switch foodType {
	case 1:
		idx := rand.Intn(len(allFood.Data.TrainFoodList))
		ctx.Set(FoodName, allFood.Data.TrainFoodList[idx].FoodName)
		ctx.Set(Price, allFood.Data.TrainFoodList[idx].Price)
	case 2:
		for _, v := range allFood.Data.FoodStoreListMap {
			if len(v) != 0 {
				idx := rand.Intn(len(v))
				ctx.Set(StoreName, v[idx].StoreName)
				ctx.Set(FoodName, v[idx].FoodList[rand.Intn(len(v[idx].FoodList))].FoodName)
				ctx.Set(Price, v[idx].FoodList[rand.Intn(len(v[idx].FoodList))].Price)
			}
		}
	}
	ctx.Set(FoodType, foodType)
	return nil, nil
}

func CreateFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("SvcImpl client not found in context")
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

	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		log.Errorf("Failed to create food order: %v | FoodOrder: %+v", err, foodOrder)
		return nil, err
	}
	if newCreateResp.Status != 1 {
		log.Errorf("Food order creation returned non-success status: %d | Response: %+v", newCreateResp.Status, newCreateResp)
		return nil, fmt.Errorf("food order creation failed, status: %d", newCreateResp.Status)
	}

	ctx.Set(OrderId, newCreateResp.Data.OrderId)
	ctx.Set(FoodType, newCreateResp.Data.FoodType)
	ctx.Set(StoreName, newCreateResp.Data.StoreName)
	ctx.Set(FoodName, newCreateResp.Data.FoodName)
	ctx.Set(Price, newCreateResp.Data.Price)

	return nil, nil
}
