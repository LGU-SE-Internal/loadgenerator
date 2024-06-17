package service

import (
	"encoding/json"
	"fmt"
	"io"
)

// TrainType represents the train type structure
//type TrainType struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//	// Add other fields as necessary
//}

// Response represents a generic response structure
//type Response struct {
//	Status int         `json:"status"`
//	Msg    string      `json:"msg"`
//	Data   interface{} `json:"data"`
//}

// TrainService defines the methods that the service should implement
type TrainService interface {
	Create(trainType TrainType) (*Response, error)
	Retrieve(id string) (*TrainType, error)
	RetrieveByName(name string) (*TrainType, error)
	RetrieveByNames(names []string) ([]TrainType, error)
	Update(trainType TrainType) (*Response, error)
	Delete(id string) (*Response, error)
	Query() ([]TrainType, error)
}

type trainCreateResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) Create(trainType *TrainType) (*trainCreateResponse, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, trainType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result trainCreateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type TrainRetrieveTrainType struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EconomyClass int    `json:"economyClass"`
		ConfortClass int    `json:"confortClass"`
		AverageSpeed int    `json:"averageSpeed"`
	} `json:"data"`
}

func (s *SvcImpl) Retrieve(id string) (*TrainRetrieveTrainType, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains/%s", s.BaseUrl, id)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TrainRetrieveTrainType
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type TrainRetrieveByNameType struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EconomyClass int    `json:"economyClass"`
		ConfortClass int    `json:"confortClass"`
		AverageSpeed int    `json:"averageSpeed"`
	} `json:"data"`
}

func (s *SvcImpl) RetrieveByName(name string) (*TrainRetrieveByNameType, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains/byName/%s", s.BaseUrl, name)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TrainRetrieveByNameType
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type TrainRetrieveByNamesType struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EconomyClass int    `json:"economyClass"`
		ConfortClass int    `json:"confortClass"`
		AverageSpeed int    `json:"averageSpeed"`
	} `json:"data"`
}

func (s *SvcImpl) RetrieveByNames(names []string) (*TrainRetrieveByNamesType, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains/byNames", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, names)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TrainRetrieveByNamesType
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type trainUpdateResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   bool   `json:"data"`
}

func (s *SvcImpl) Update(trainType *TrainType) (*trainUpdateResponse, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains", s.BaseUrl)
	resp, err := s.cli.SendRequest("PUT", url, trainType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result trainUpdateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type trainDeleteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   bool   `json:"data"`
}

func (s *SvcImpl) Delete(id string) (*trainDeleteResponse, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains/%s", s.BaseUrl, id)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result trainDeleteResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type TrainResponseType struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EconomyClass int    `json:"economyClass"`
		ConfortClass int    `json:"confortClass"`
		AverageSpeed int    `json:"averageSpeed"`
	} `json:"data"`
}

func (s *SvcImpl) Query() (*TrainResponseType, error) {
	url := fmt.Sprintf("%s/api/v1/trainservice/trains", s.BaseUrl)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TrainResponseType
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
