package service

import (
	"encoding/json"
	"io"
)

type FoodService interface {
	FindAllFoodOrder() (*FindAllFoodOrder, error)
	CreateFoodOrder(foodOrder *FoodOrder) (*CreateFoodOrderResp, error)
	CreateFoodOrdersInBatch(foodOrders []FoodOrder) (*CreateFoodOrdersInBatch, error)
	UpdateFoodOrder(foodOrder *FoodOrder) (*FoodOrder, error)
	DeleteFoodOrder(orderID string) (*DeleteFoodOrderResp, error)
	FindByOrderId(orderID string) (*FindByOrderIdResponse, error)
	GetAllFood(date, startStation, endStation, tripID string) (*GetAllFoodResponse, error)
}

// FoodOrder represents the food order structure
type FoodOrder struct {
	ID          string  `json:"id"`
	OrderID     string  `json:"orderId"`
	FoodType    int     `json:"foodType"` // 1: train food; 2: food store
	StationName string  `json:"stationName"`
	StoreName   string  `json:"storeName"`
	FoodName    string  `json:"foodName"`
	Price       float64 `json:"price"`
}

type CreateFoodOrderResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id          string      `json:"id"`
		OrderId     string      `json:"orderId"`
		FoodType    int         `json:"foodType"`
		StationName interface{} `json:"stationName"`
		StoreName   interface{} `json:"storeName"`
		FoodName    string      `json:"foodName"`
		Price       float64     `json:"price"`
	} `json:"data"`
}

type DeleteFoodOrderResp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type FindAllFoodOrder struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id          string  `json:"id"`
		OrderId     string  `json:"orderId"`
		FoodType    int     `json:"foodType"`
		StationName string  `json:"stationName"`
		StoreName   string  `json:"storeName"`
		FoodName    string  `json:"foodName"`
		Price       float64 `json:"price"`
	} `json:"data"`
}

func (s *SvcImpl) FindAllFoodOrder() (*FindAllFoodOrder, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/foodservice/orders", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result FindAllFoodOrder
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) CreateFoodOrder(foodOrder *FoodOrder) (*CreateFoodOrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/foodservice/orders", foodOrder)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CreateFoodOrderResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateFoodOrdersInBatch struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) CreateFoodOrdersInBatch(foodOrders []FoodOrder) (*CreateFoodOrdersInBatch, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/foodservice/createOrderBatch", foodOrders)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CreateFoodOrdersInBatch
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) UpdateFoodOrder(foodOrder *FoodOrder) (*FoodOrder, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/foodservice/orders", foodOrder)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result FoodOrder
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) DeleteFoodOrder(orderID string) (*DeleteFoodOrderResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+"/api/v1/foodservice/orders/"+orderID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result DeleteFoodOrderResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type FindByOrderIdResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id          string  `json:"id"`
		OrderId     string  `json:"orderId"`
		FoodType    int     `json:"foodType"`
		StationName string  `json:"stationName"`
		StoreName   string  `json:"storeName"`
		FoodName    string  `json:"foodName"`
		Price       float64 `json:"price"`
	} `json:"data"`
}

func (s *SvcImpl) FindByOrderId(orderID string) (*FindByOrderIdResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/foodservice/orders/"+orderID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result FindByOrderIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetAllFoodResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) GetAllFood(date string, startStation string, endStation string, tripID string) (*GetAllFoodResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/foodservice/foods/"+date+"/"+startStation+"/"+endStation+"/"+tripID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetAllFoodResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
