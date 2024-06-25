package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SecurityService defines the interface for security operations
type SecurityService interface {
	FindAllSecurityConfig() (*FindAllResponse, error)
	AddNewSecurityConfig(config *SecurityConfig) (*SingleResponse, error)
	ModifySecurityConfig(config *SecurityConfig) (*SingleResponse, error)
	DeleteSecurityConfig(id string) (*DeleteResponse, error)
	Check(accountId string) (*SingleResponse, error)
}
type SecurityConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

// FindAllResponse represents the response structure for finding all security configs
type FindAllResponse struct {
	Status int              `json:"status"`
	Msg    string           `json:"msg"`
	Data   []SecurityConfig `json:"data"`
}

// SingleResponse represents the response structure for single security config operations
type SingleResponse struct {
	Status int            `json:"status"`
	Msg    string         `json:"msg"`
	Data   SecurityConfig `json:"data"`
}

// DeleteResponse represents the response structure for delete operations
//type DeleteResponse struct {
//	Status int         `json:"status"`
//	Msg    string      `json:"msg"`
//	Data   interface{} `json:"data"`
//}

func (s *SvcImpl) FindAllSecurityConfig() (*FindAllResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/securityservice/securityConfigs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result FindAllResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) AddNewSecurityConfig(config *SecurityConfig) (*SingleResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/securityservice/securityConfigs", config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SingleResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) ModifySecurityConfig(config *SecurityConfig) (*SingleResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/securityservice/securityConfigs", config)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SingleResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) DeleteSecurityConfig(id string) (*DeleteResponse, error) {
	url := fmt.Sprintf("%s/api/v1/securityservice/securityConfigs/%s", s.BaseUrl, id)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result DeleteResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) Check(accountId string) (*SingleResponse, error) {
	url := fmt.Sprintf("%s/api/v1/securityservice/securityConfigs/%s", s.BaseUrl, accountId)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SingleResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}
