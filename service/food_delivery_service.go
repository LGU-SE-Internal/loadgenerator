package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqCreateFoodDeliveryOrder(input *FoodDeliveryOrder) (*FoodDeliveryOrderResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/fooddeliveryservice/orders", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetAllFoodDeliveryOrders() (*FoodDeliveryOrderArrResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/all", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetFoodDeliveryOrderByStoreId(storeId string) (*FoodDeliveryOrderArrResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/store/"+storeId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetFoodDeliveryOrderById(orderId string) (*FoodDeliveryOrderResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqDeleteFoodDeliveryOrderById(orderId string) (*DataStringResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/d/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DataStringResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateDeliveryTime(input *DeliveryInfo) (*FoodDeliveryOrderResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/dtime", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateSeatNo(input *SeatInfo) (*FoodDeliveryOrderResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/seatno", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateTripId(input *TripOrderInfo) (*FoodDeliveryOrderResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/fooddeliveryservice/orders/tripid", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result FoodDeliveryOrderResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
