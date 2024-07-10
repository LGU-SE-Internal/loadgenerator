package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// TrainType represents the train type structure
//type TrainType struct {
//	ID           string `json:"id"`
//	Name         string `json:"name"`
//	EconomyClass int    `json:"economyClass"`
//	ConfortClass int    `json:"confortClass"`
//	AverageSpeed int    `json:"averageSpeed"`
//}

type TrainType struct {
	ID           string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;size:36" json:"id"`
	Name         string `gorm:"not null;unique" json:"name"`
	EconomyClass int    `gorm:"column:economy_class" json:"economy_class"`
	ConfortClass int    `gorm:"column:confort_class" json:"confort_class"`
	AverageSpeed int    `gorm:"column:average_speed" json:"average_speed"`
}

// DeleteStationResponse represents a generic response structure
//type DeleteStationResponse struct {
//	Status int         `json:"status"`
//	Msg    string      `json:"msg"`
//	Data   interface{} `json:"data"`
//}

// TrainService defines the methods that the service should implement
type TrainService interface {
	Create(trainType *TrainType) (*CreateStationResponse, error)
	Retrieve(id string) (*TrainServiceRetrieveTrainType, error)
	RetrieveByName(name string) (*TrainRetrieveByNameType, error)
	RetrieveByNames(names []string) (*TrainRetrieveByNamesType, error)
	Update(trainType *TrainType) (*TrainUpdateResponse, error)
	Delete(id string) (*TrainDeleteResponse, error)
	Query() (*TrainResponseType, error)
}

type CreateStationResponse struct {
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

func (s *SvcImpl) Create(trainType *TrainType) (*CreateStationResponse, error) {
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

	var result CreateStationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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

type TrainServiceRetrieveTrainType struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) Retrieve(id string) (*TrainServiceRetrieveTrainType, error) {
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

	var result TrainServiceRetrieveTrainType
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type TrainUpdateResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   bool   `json:"data"`
}

func (s *SvcImpl) Update(trainType *TrainType) (*TrainUpdateResponse, error) {
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

	var result TrainUpdateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type TrainDeleteResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   bool   `json:"data"`
}

func (s *SvcImpl) Delete(id string) (*TrainDeleteResponse, error) {
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

	var result TrainDeleteResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}
