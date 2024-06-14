package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type ConsignPrice struct {
	ID            string  `json:"id"`             // `@Id @GeneratedValue(generator = "jpa-uuid") @Column(length = 36)`
	Index         int     `json:"index"`          // `@Column(name = "idx", unique = true)`
	InitialWeight float64 `json:"initial_weight"` // `@Column(name = "initial_weight")`
	InitialPrice  float64 `json:"initial_price"`  // `@Column(name = "initial_price")`
	WithinPrice   float64 `json:"within_price"`   // `@Column(name = "within_price")`
	BeyondPrice   float64 `json:"beyond_price"`   // `@Column(name = "beyond_price")`
}

func (s *SvcImpl) GetPriceByWeightAndRegion(weight string, isWithinRegion string) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) GetPriceInfo() (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) GetPriceConfig() (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) ModifyPriceConfig(priceConfig ConsignPrice) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
