package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type PriceService interface {
	FindByRouteIdAndTrainType(routeId, trainType string) (*AdminPriceResponse, error)
	FindByRouteIdsAndTrainTypes(ridsAndTts []string) (*AllPriceResponse, error)
	FindAllPriceConfig() (*AllPriceResponse, error)
	CreateNewPriceConfig(info PriceConfig) (*AdminPriceResponse, error)
	DeletePriceConfig(pricesId string) (*AdminPriceResponse, error)
	UpdatePriceConfig(info PriceConfig) (*AdminPriceResponse, error)
}

type PriceConfig struct {
	Id        string  `json:"id"`
	RouteId   string  `json:"routeId"`
	TrainType string  `json:"trainType"`
	Price     float64 `json:"price"`
	// Add more fields as needed
}

//type AdminPriceResponse struct {
//	Status int         `json:"status"`
//	Msg    string      `json:"msg"`
//	Data   PriceConfig `json:"data"`
//}

type AllPriceResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []PriceConfig `json:"data"`
}

func (s *SvcImpl) FindByRouteIdAndTrainType(routeId, trainType string) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/priceservice/prices/%s/%s", routeId, trainType), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) FindByRouteIdsAndTrainTypes(ridsAndTts []string) (*AllPriceResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/priceservice/prices/byRouteIdsAndTrainTypes", ridsAndTts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AllPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) FindAllPriceConfig() (*AllPriceResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/priceservice/prices", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AllPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) CreateNewPriceConfig(info PriceConfig) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/priceservice/prices", info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeletePriceConfig(pricesId string) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/priceservice/prices/%s", pricesId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) UpdatePriceConfig(info PriceConfig) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/priceservice/prices", info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
