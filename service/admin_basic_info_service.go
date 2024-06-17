package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type GetContactsResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id             string  `json:"id"`
		AccountId      *string `json:"accountId"`
		Name           string  `json:"name"`
		DocumentType   int     `json:"documentType"`
		DocumentNumber string  `json:"documentNumber"`
		PhoneNumber    string  `json:"phoneNumber"`
	} `json:"data"`
}

// RouteResponse structs
type ContactResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id             string `json:"id"`
		AccountId      string `json:"accountId"`
		Name           string `json:"name"`
		DocumentType   int    `json:"documentType"`
		DocumentNumber string `json:"documentNumber"`
		PhoneNumber    string `json:"phoneNumber"`
	} `json:"data"`
}

type StationResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type deleteResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type TrainResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type ConfigResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type PriceResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id                  string  `json:"id"`
		TrainType           string  `json:"trainType"`
		RouteId             string  `json:"routeId"`
		BasicPriceRate      float64 `json:"basicPriceRate"`
		FirstClassPriceRate float64 `json:"firstClassPriceRate"`
	} `json:"data"`
}

// Request structs
type Contacts struct {
	ID             string `json:"id"`
	AccountID      string `json:"accountId"`
	Name           string `json:"name"`
	DocumentType   int    `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	PhoneNumber    string `json:"phoneNumber"`
}

type Station struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	StayTime int    `json:"stayTime"`
}

type TrainType struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	EconomyClass int    `json:"economyClass"`
	ConfortClass int    `json:"confortClass"`
	AverageSpeed int    `json:"averageSpeed"`
}

type Config struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
type ModifyContactsResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id             string `json:"id"`
		AccountId      string `json:"accountId"`
		Name           string `json:"name"`
		DocumentType   int    `json:"documentType"`
		DocumentNumber string `json:"documentNumber"`
		PhoneNumber    string `json:"phoneNumber"`
	} `json:"data"`
}
type PriceInfo struct {
	ID                  string  `json:"id"`
	TrainType           string  `json:"trainType"`
	RouteID             string  `json:"routeId"`
	BasicPriceRate      float64 `json:"basicPriceRate"`
	FirstClassPriceRate float64 `json:"firstClassPriceRate"`
}

// AdminBasicInfoService methods

func (s *SvcImpl) AdminGetAllContacts() (*GetContactsResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/contacts", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetContactsResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

type adminDeleteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) AdminDeleteContact(contactsId string) (*adminDeleteResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/contacts/%s", contactsId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result adminDeleteResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminModifyContact(contacts *Contacts) (*ContactResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/contacts", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminAddContact(contacts *Contacts) (*ContactResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/contacts", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) GetAllStations() (*StationResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/stations", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result StationResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) DeleteStation(id string) (*deleteResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/stations/%s", id), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result deleteResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) ModifyStation(station *Station) (*StationResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/stations", station)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result StationResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AddStation(station *Station) (*StationResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/stations", station)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result StationResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) GetAllTrains() (*TrainResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/trains", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) DeleteTrain(id string) (*TrainResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/trains/%s", id), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) ModifyTrain(train *TrainType) (*TrainResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/trains", train)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AddTrain(train *TrainType) (*TrainResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/trains", train)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) GetAllConfigs() (*ConfigResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/configs", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) DeleteConfig(name string) (*ConfigResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/configs/%s", name), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) ModifyConfig(config *Config) (*ConfigResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/configs", config)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AddConfig(config *Config) (*ConfigResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/configs", config)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) GetAllPrices() (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/prices", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) DeletePrice(pricesId string) (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/prices/%s", pricesId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) ModifyPrice(price *PriceInfo) (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/prices", price)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AddPrice(price *PriceInfo) (*PriceResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/prices", price)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}
