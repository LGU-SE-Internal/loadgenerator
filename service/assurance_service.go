package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type AssuranceService interface {
	GetAllAssurances() ([]AssuranceInfo, error)
	GetAllAssuranceTypes() ([]AssuranceType, error)
	DeleteAssuranceByID(assuranceID string) (*AssuranceResponse, error)
	DeleteAssuranceByOrderID(orderID string) (*AssuranceResponse, error)
	ModifyAssurance(assuranceID, orderID string, typeIndex int) (*AssuranceResponse, error)
	CreateNewAssurance(typeIndex int, orderID string) (*AssuranceResponse, error)
	GetAssuranceByID(assuranceID string) (*AssuranceInfo, error)
	FindAssuranceByOrderID(orderID string) ([]AssuranceInfo, error)
}

type AssuranceInfo struct {
	// Define the structure based on your API response
	AssuranceID string `json:"assuranceId"`
	OrderID     string `json:"orderId"`
	// Add more fields as needed
}

type AssuranceType struct {
	// Define the structure based on your API response
	TypeID   int    `json:"typeId"`
	TypeName string `json:"typeName"`
	// Add more fields as needed
}

type AssuranceResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		AssuranceID string `json:"assuranceId"`
	} `json:"data"`
}

func (s *SvcImpl) GetAllAssurances() ([]AssuranceInfo, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/assuranceservice/assurances", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assurances []AssuranceInfo
	err = json.NewDecoder(resp.Body).Decode(&assurances)
	if err != nil {
		return nil, err
	}

	return assurances, nil
}

func (s *SvcImpl) GetAllAssuranceTypes() ([]AssuranceType, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/assuranceservice/assurances/types", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var types []AssuranceType
	err = json.NewDecoder(resp.Body).Decode(&types)
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (s *SvcImpl) DeleteAssuranceByID(assuranceID string) (*AssuranceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/assuranceid/%s", s.BaseUrl, assuranceID)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeleteAssuranceByOrderID(orderID string) (*AssuranceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/orderid/%s", s.BaseUrl, orderID)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) ModifyAssurance(assuranceID string, orderID string, typeIndex int) (*AssuranceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/%s/%s/%d", s.BaseUrl, assuranceID, orderID, typeIndex)
	resp, err := s.cli.SendRequest("PATCH", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) CreateNewAssurance(typeIndex int, orderID string) (*AssuranceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/%d/%s", s.BaseUrl, typeIndex, orderID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AssuranceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetAssuranceByID(assuranceID string) (*AssuranceInfo, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/assuranceid/%s", s.BaseUrl, assuranceID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assurance AssuranceInfo
	err = json.NewDecoder(resp.Body).Decode(&assurance)
	if err != nil {
		return nil, err
	}

	return &assurance, nil
}

func (s *SvcImpl) FindAssuranceByOrderID(orderID string) ([]AssuranceInfo, error) {
	url := fmt.Sprintf("%s/api/v1/assuranceservice/assurances/orderid/%s", s.BaseUrl, orderID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var assurances []AssuranceInfo
	err = json.NewDecoder(resp.Body).Decode(&assurances)
	if err != nil {
		return nil, err
	}

	return assurances, nil
}
