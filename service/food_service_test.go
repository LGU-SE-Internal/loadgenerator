package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_FoodService(t *testing.T) {
	cli, _ := GetBasicClient()

	// Test CreateFoodOrder
	createResp, err := cli.CreateFoodOrder(&FoodOrder{
		OrderID:     faker.UUIDHyphenated(),
		FoodName:    "HotPot",
		StationName: "Shang Hai",
		StoreName:   "MiaoTing Instant-Boiled Mutton",
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("create response: %+v", createResp)

	// Test FindAllFoodOrder
	resp, err := cli.FindAllFoodOrder()
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if len(resp) == 0 {
		t.Errorf("FindAllFoodOrder returned no results")
	}

	// Mock data
	MockedOrderID := faker.UUIDHyphenated()
	foodOrder := FoodOrder{
		OrderID:     MockedOrderID,
		FoodName:    "HotPot",
		StationName: "Shang Hai",
		StoreName:   "MiaoTing Instant-Boiled Mutton",
	}

	// Create Test
	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		t.Errorf("NewCreateFoodOrder request failed, err %s", err)
	}
	if newCreateResp.Status != 1 {
		t.Errorf("NEwCreateFoodOrder failed")
	}

	// Query all
	allFoodOrders, err := cli.FindAllFoodOrder()
	if err != nil {
		t.Errorf("FindAllFoodOrder request failed, err %s", err)
	}
	if len(allFoodOrders) == 0 {
		t.Errorf("FindAllFoodOrder returned no results")
	}

	var getOrderId string
	if len(allFoodOrders) > 0 {
		getOrderId = allFoodOrders[0].OrderID
	}

	// Test Update
	updateFoodOrder := FoodOrder{
		OrderID:     getOrderId,
		FoodName:    "HotPot",
		StationName: "Shang Hai",
		StoreName:   "MiaoTing Instant-Boiled Mutton",
	}
	updateResp, err := cli.UpdateFoodOrder(&updateFoodOrder)
	if err != nil {
		t.Errorf("UpdateFoodOrder request failed, err %s", err)
	}
	if updateResp.OrderID == "" {
		t.Errorf("UpdateFoodOrder failed")
	}

	// Test Delete
	var deleteOrderID string
	if len(allFoodOrders) > 0 {
		deleteOrderID = allFoodOrders[len(allFoodOrders)-1].OrderID
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
	if len(allFoodResp) == 0 {
		t.Errorf("GetAllFood returned no results")
	}

	// Test CreateFoodOrdersInBatch
	foodOrders := []FoodOrder{foodOrder, updateFoodOrder}
	createBatchResp, err := cli.CreateFoodOrdersInBatch(foodOrders)
	if err != nil {
		t.Errorf("CreateFoodOrdersInBatch request failed, err %s", err)
	}
	if len(createBatchResp) == 0 {
		t.Errorf("CreateFoodOrdersInBatch failed")
	}

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
	if len(findByTrainNumberResp) == 0 {
		t.Errorf("GetAllFood by train number returned no results")
	}
}
