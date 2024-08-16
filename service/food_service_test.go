package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_FoodService(t *testing.T) {
	cli, _ := GetBasicClient()

	CreateInpt := &FoodOrder{
		ID:          faker.UUIDHyphenated(),
		OrderID:     faker.UUIDHyphenated(),
		FoodType:    1,
		FoodName:    "HotPot",
		StationName: "Shang Hai",
		StoreName:   "MiaoTing Instant-Boiled Mutton",
		Price:       7.00,
	}
	// Test CreateFoodOrder
	createResp, err := cli.CreateFoodOrder(CreateInpt)
	if err != nil {
		t.Error(err)
	}
	t.Logf("create response: %+v", createResp)

	// Test FindAllFoodOrder
	resp, err := cli.FindAllFoodOrder()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if len(resp.Data) == 0 {
		t.Errorf("FindAllFoodOrder returned no results")
	}

	// Mock data
	MockedOrderID := faker.UUIDHyphenated()
	MockedID := faker.UUIDHyphenated()
	foodOrder := FoodOrder{
		ID:          MockedID,
		OrderID:     MockedOrderID,
		FoodType:    1,
		FoodName:    "HotPot",
		StationName: "Shang Hai",
		StoreName:   "MiaoTing Instant-Boiled Mutton",
		Price:       7.00,
	}

	// Create Test
	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		t.Errorf("NewCreateFoodOrder request failed, err %s", err)
	}
	if newCreateResp.Status != 1 {
		t.Errorf("NEwCreateFoodOrder failed")
	}

	// QueryTraintype all
	allFoodOrders, err := cli.FindAllFoodOrder()
	if err != nil {
		t.Errorf("FindAllFoodOrder request failed, err %s", err)
	}
	if len(allFoodOrders.Data) == 0 {
		t.Errorf("FindAllFoodOrder returned no results")
	}

	var getOrderId string
	//var getId string
	if len(allFoodOrders.Data) > 0 {
		getOrderId = allFoodOrders.Data[0].OrderId
		//getId = allFoodOrders.Data[0].Id
	}

	// Test Update
	updateFoodOrder := FoodOrder{
		ID:          MockedID,
		OrderID:     MockedOrderID,
		FoodType:    1,
		FoodName:    "HotPot",
		StationName: "Shang Hai",
		StoreName:   "MiaoTing Instant-Boiled Mutton",
		Price:       8.00,
	}
	updateResp, err := cli.UpdateFoodOrder(&updateFoodOrder)
	if err != nil {
		t.Errorf("UpdateFoodOrder request failed, err %s", err)
	}
	t.Logf("UpdateFoodOrder return: %v", updateResp)

	// Test Delete
	var deleteOrderID string
	if len(allFoodOrders.Data) > 0 {
		deleteOrderID = allFoodOrders.Data[len(allFoodOrders.Data)-1].OrderId
	} else {
		t.Errorf("FindAllFoodOrder returned empty data")
	}
	deleteResp, err := cli.DeleteFoodOrder(deleteOrderID)
	if err != nil {
		t.Errorf("DeleteFoodOrder request failed, err %s", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("DeleteFoodOrder failed")
	}

	// Test FindByOrderId
	findByOrderIDResp, err := cli.FindByOrderId(getOrderId)
	if err != nil {
		t.Errorf("FindByOrderId request failed, err %s", err)
	}
	if findByOrderIDResp == nil {
		t.Errorf("FindByOrderId returned no result")
	}

	// Test GetAllFood
	allFoodResp, err := cli.GetAllFood("2024-06-06", "Shang Hai", "Beijing", "trip123")
	if err != nil {
		t.Errorf("GetAllFood request failed, err %s", err)
	}
	t.Logf("GetAllFood returned results： %v", allFoodResp)

	// Test CreateFoodOrdersInBatch
	foodOrders := []FoodOrder{foodOrder, updateFoodOrder}
	createBatchResp, err := cli.CreateFoodOrdersInBatch(foodOrders)
	if err != nil {
		t.Errorf("CreateFoodOrdersInBatch request failed, err %s", err)
	}
	t.Logf("CreateFoodOrdersInBatch failed: %v", createBatchResp)

	// Test finding by random train number
	trainNumber := ""
	randn := rand.Int() % 15
	if randn > 7 {
		trainNumber = "G" + strconv.Itoa(rand.Int()%100)
	} else if randn > 3 {
		trainNumber = "D" + strconv.Itoa(rand.Int()%100)
	} else {
		trainNumber = strconv.Itoa(rand.Int() % 100)
	}
	t.Logf("trainNumber: %s", trainNumber)

	findByTrainNumberResp, err := cli.GetAllFood("2024-06-06", "Shang Hai", "Beijing", trainNumber)
	if err != nil {
		t.Errorf("GetAllFood request failed, err %s", err)
	}
	t.Logf("GetAllFood by train number returned results: %v", findByTrainNumberResp)

}
