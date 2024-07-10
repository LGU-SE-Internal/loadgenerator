package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type AdminBasicInfoService interface {
	AdminGetAllContacts() (*AdminGetContactsResp, error)
	AdminDeleteContact(contactsId string) (*AdminDeleteContactResp, error)
	AdminModifyContact(contacts *AdminContacts) (*AdminContactResponse, error)
	AdminAddContact(contacts *AdminContacts) (*AdminContactResponse, error)
	AdminGetAllStations() (*AdminStationResponse, error)
	AdminDeleteStation(id string) (*AdminDeleteResponse, error)
	AdminModifyStation(station *AdminStation) (*AdminStationResponse, error)
	AdminAddStation(station *AdminStation) (*AdminStationResponse, error)
	AdminGetAllTrains() (*AdminTrainResponse, error)
	AdminDeleteTrain(id string) (*AdminTrainResponse, error)
	AdminModifyTrain(train *AdminTrainType) (*AdminTrainResponse, error)
	AdminAddTrain(train *AdminTrainType) (*AdminTrainResponse, error)
	AdminGetAllConfigs() (*AdminConfigResponse, error)
	AdminDeleteConfig(name string) (*AdminConfigResponse, error)
	AdminModifyConfig(config *AdminConfig) (*AdminConfigResponse, error)
	AdminAddConfig(config *AdminConfig) (*AdminConfigResponse, error)
	AdminGetAllPrices() (*AdminPriceResponse, error)
	AdminDeletePrice(pricesId string) (*AdminPriceResponse, error)
	AdminModifyPrice(price *AdminPriceInfo) (*AdminPriceResponse, error)
	AdminAddPrice(price *AdminPriceInfo) (*AdminPriceResponse, error)
}
type AdminGetContactsResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id             string `json:"id"`
		AccountId      string `json:"accountId"`
		Name           string `json:"name"`
		DocumentType   int    `json:"documentType"`
		DocumentNumber string `json:"documentNumber"`
		PhoneNumber    string `json:"phoneNumber"`
	} `json:"data"`

}

type AdminContactResponse struct {
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
type AdminDeleteContactResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}
type AdminStationResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type AdminDeleteResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type AdminTrainResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type AdminConfigResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type AdminPriceResponse struct {
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
type AdminContacts struct {
	ID             string `json:"id"`
	AccountID      string `json:"accountId"`
	Name           string `json:"name"`
	DocumentType   int    `json:"documentType"`
	DocumentNumber string `json:"documentNumber"`
	PhoneNumber    string `json:"phoneNumber"`
}

type AdminStation struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	StayTime int    `json:"stayTime"`
}

type AdminTrainType struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	EconomyClass int    `json:"economyClass"`
	ConfortClass int    `json:"confortClass"`
	AverageSpeed int    `json:"averageSpeed"`
}

type AdminConfig struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
type AdminModifyContactsResp struct {
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
type AdminPriceInfo struct {
	ID                  string  `json:"id"`
	TrainType           string  `json:"trainType"`
	RouteID             string  `json:"routeId"`
	BasicPriceRate      float64 `json:"basicPriceRate"`
	FirstClassPriceRate float64 `json:"firstClassPriceRate"`
}

// AdminBasicInfoService methods

func (s *SvcImpl) AdminGetAllContacts() (*AdminGetContactsResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/contacts", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminGetContactsResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminDeleteContact(contactsId string) (*AdminDeleteContactResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/contacts/%s", contactsId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminDeleteContactResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminModifyContact(contacts *AdminContacts) (*AdminContactResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/contacts", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminAddContact(contacts *AdminContacts) (*AdminContactResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/contacts", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminGetAllStations() (*AdminStationResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/stations", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminStationResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminDeleteStation(id string) (*AdminDeleteResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/stations/%s", id), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminDeleteResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminModifyStation(station *AdminStation) (*AdminStationResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/stations", station)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminStationResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminAddStation(station *AdminStation) (*AdminStationResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/stations", station)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminStationResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminGetAllTrains() (*AdminTrainResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/trains", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminTrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminDeleteTrain(id string) (*AdminTrainResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/trains/%s", id), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminTrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminModifyTrain(train *AdminTrainType) (*AdminTrainResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/trains", train)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminTrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminAddTrain(train *AdminTrainType) (*AdminTrainResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/trains", train)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminTrainResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminGetAllConfigs() (*AdminConfigResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/configs", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminDeleteConfig(name string) (*AdminConfigResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/configs/%s", name), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminModifyConfig(config *AdminConfig) (*AdminConfigResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/configs", config)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminAddConfig(config *AdminConfig) (*AdminConfigResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/configs", config)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminConfigResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminGetAllPrices() (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/prices", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminDeletePrice(pricesId string) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/adminbasicservice/adminbasic/prices/%s", pricesId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminModifyPrice(price *AdminPriceInfo) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/prices", price)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AdminAddPrice(price *AdminPriceInfo) (*AdminPriceResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminbasicservice/adminbasic/prices", price)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminPriceResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}
