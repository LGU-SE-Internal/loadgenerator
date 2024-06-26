package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type ConsignPriceService interface {
	GetPriceByWeightAndRegion(weight string, isWithinRegion string) (*ConsignPriceResponse, error)
	GetPriceInfo() (*GetResponse, error)
	GetPriceConfig() (*GetPriceConfigResponse, error)
	ModifyPriceConfig(priceConfig *ConsignPrice) (*ModifyConsignPriceResponse, error)
}
type ConsignPrice struct {
	ID            string  `json:"id"`            // UUID field, mapped from Java's String
	Index         int     `json:"index"`         // Unique index field
	InitialWeight float64 `json:"initialWeight"` // Initial weight field
	InitialPrice  float64 `json:"initialPrice"`  // Initial price field
	WithinPrice   float64 `json:"withinPrice"`   // Within price field
	BeyondPrice   float64 `json:"beyondPrice"`   // Beyond price field
}

type ConsignPriceResponse struct {
	Status int     `json:"status"`
	Msg    string  `json:"msg"`
	Data   float64 `json:"data"`
}

type GetResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type GetPriceConfigResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id            string  `json:"id"`
		Index         int     `json:"index"`
		InitialWeight float64 `json:"initialWeight"`
		InitialPrice  float64 `json:"initialPrice"`
		WithinPrice   float64 `json:"withinPrice"`
		BeyondPrice   float64 `json:"beyondPrice"`
	} `json:"data"`
}
type ModifyConsignPriceResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id            string  `json:"id"`
		Index         int     `json:"index"`
		InitialWeight float64 `json:"initialWeight"`
		InitialPrice  float64 `json:"initialPrice"`
		WithinPrice   float64 `json:"withinPrice"`
		BeyondPrice   float64 `json:"beyondPrice"`
	} `json:"data"`
}

func (s *SvcImpl) GetPriceByWeightAndRegion(weight string, isWithinRegion string) (*ConsignPriceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/consignpriceservice/consignprice/%s/%s", s.BaseUrl, weight, isWithinRegion)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ConsignPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) GetPriceInfo() (*GetResponse, error) {
	url := fmt.Sprintf("%s/api/v1/consignpriceservice/consignprice/price", s.BaseUrl)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) GetPriceConfig() (*GetPriceConfigResponse, error) {
	url := fmt.Sprintf("%s/api/v1/consignpriceservice/consignprice/config", s.BaseUrl)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetPriceConfigResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) ModifyPriceConfig(priceConfig *ConsignPrice) (*ModifyConsignPriceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/consignpriceservice/consignprice", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, priceConfig)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ModifyConsignPriceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}
