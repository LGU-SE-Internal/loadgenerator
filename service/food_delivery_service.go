package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqCreateFoodDeliveryOrder(input *FoodDeliveryOrder) (*DataStringResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/fooddeliveryservice/orders", input)
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
