package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type PriceConfig struct {
	ID                  string  `json:"id"`                  // id主键改成uuid类型的 自定义生成策略
	TrainType           string  `json:"trainType"`           // 这次托运关联订单
	RouteID             string  `json:"routeId"`             // 这次托运关联的账户
	BasicPriceRate      float64 `json:"basicPriceRate"`      // 基础价格
	FirstClassPriceRate float64 `json:"firstClassPriceRate"` // 一等票价格
}

//type PriceResponse struct {
//	Status int         `json:"status"`
//	Msg    string      `json:"msg"`
//	Data   PriceConfig `json:"data"`
//}

type AllPriceResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id                  string  `json:"id"`
		TrainType           string  `json:"trainType"`
		RouteId             string  `json:"routeId"`
		BasicPriceRate      float64 `json:"basicPriceRate"`
		FirstClassPriceRate float64 `json:"firstClassPriceRate"`
	} `json:"data"`
}

type PriceService interface {
	FindByRouteIdAndTrainType(routeId, trainType string) (*PriceResponse, error)
	FindByRouteIdsAndTrainTypes(ridsAndTts []string) (*AllPriceResponse, error)
	FindAllPriceConfig() (*AllPriceResponse, error)
	CreateNewPriceConfig(info PriceConfig) (*PriceResponse, error)
	DeletePriceConfig(pricesId string) (*PriceResponse, error)
	UpdatePriceConfig(info PriceConfig) (*PriceResponse, error)
}

type findByRouteIdAndTrainTypeResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) FindByRouteIdAndTrainType(routeId string, trainType string) (*findByRouteIdAndTrainTypeResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/priceservice/prices/%s/%s", routeId, trainType), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result findByRouteIdAndTrainTypeResponse
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

func (s *SvcImpl) CreateNewPriceConfig(info *PriceConfig) (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/priceservice/prices", info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeletePriceConfig(pricesId string) (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/priceservice/prices/%s", pricesId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) UpdatePriceConfig(info PriceConfig) (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/priceservice/prices", info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
