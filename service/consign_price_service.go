package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type ConsignPrice struct {
	ID            string  `json:"id"`            // UUID field, mapped from Java's String
	Index         int     `json:"index"`         // Unique index field
	InitialWeight float64 `json:"initialWeight"` // Initial weight field
	InitialPrice  float64 `json:"initialPrice"`  // Initial price field
	WithinPrice   float64 `json:"withinPrice"`   // Within price field
	BeyondPrice   float64 `json:"beyondPrice"`   // Beyond price field
}

type consignResponse struct {
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

func (s *SvcImpl) GetPriceByWeightAndRegion(weight string, isWithinRegion string) (*consignResponse, error) {
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

	var result consignResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}

	return &result, nil
}

type modifyResponse struct {
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

func (s *SvcImpl) ModifyPriceConfig(priceConfig *ConsignPrice) (*modifyResponse, error) {
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

	var result modifyResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
