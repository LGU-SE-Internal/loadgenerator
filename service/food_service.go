package service

import (
	"encoding/json"
	"io"
)

// FoodOrder represents the food order structure
type FoodOrder struct {
	OrderID     string `json:"orderId"`
	FoodName    string `json:"foodName"`
	StationName string `json:"stationName"`
	StoreName   string `json:"storeName"`
}

// CreateFoodOrderResp represents the response structure for creating a food order
type CreateFoodOrderResp struct {
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	Data   FoodOrder `json:"data"`
}

// DeleteFoodOrderResp represents the response structure for deleting a food order
type DeleteFoodOrderResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) FindAllFoodOrder() ([]FoodOrder, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/foodservice/orders", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []FoodOrder
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
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

func (s *SvcImpl) CreateFoodOrdersInBatch(foodOrders []FoodOrder) ([]FoodOrder, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/foodservice/createOrderBatch", foodOrders)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []FoodOrder
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
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

func (s *SvcImpl) FindByOrderId(orderID string) (*FoodOrder, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/foodservice/orders/"+orderID, nil)
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

func (s *SvcImpl) GetAllFood(date, startStation, endStation, tripID string) ([]FoodOrder, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/foodservice/foods/"+date+"/"+startStation+"/"+endStation+"/"+tripID, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []FoodOrder
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
