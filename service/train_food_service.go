package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type GetTrainFoodResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id       string `json:"id"`
		TripId   string `json:"tripId"`
		FoodList []struct {
			FoodName string  `json:"foodName"`
			Price    float64 `json:"price"`
		} `json:"foodList"`
	} `json:"data"`
}

type GetTrainFoodByIdResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		FoodName string  `json:"foodName"`
		Price    float64 `json:"price"`
	} `json:"data"`
}

func (s *SvcImpl) GetAllTrainFood() (*GetTrainFoodResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/trainfoodservice/trainfoods", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetTrainFoodResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) GetTrainFoodByTripId(tripId string) (*GetTrainFoodByIdResp, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/trainfoodservice/trainfoods/%s", s.BaseUrl, tripId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetTrainFoodByIdResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
