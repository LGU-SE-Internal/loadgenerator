package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type AdminTravelService interface {
	CreateTravel(request *AdminTravelInfo) (*AdminTravelResponse, error)
	UpdateTravel(request *AdminTravelInfo) (*AdminTravelResponse, error)
	DeleteTravel(tripId string) (*AdminTravelResponse, error)
	GetAllTravels() ([]AdminTravelInfo, error)
}
type AdminTravelInfo struct {
	LoginID             string `json:"loginId"`
	TripID              string `json:"tripId"`
	TrainTypeName       string `json:"trainTypeName"`
	RouteID             string `json:"routeId"`
	StartStationName    string `json:"startStationName"`
	StationsName        string `json:"stationsName"`
	TerminalStationName string `json:"terminalStationName"`
	StartTime           string `json:"startTime"`
	EndTime             string `json:"endTime"`
}

type AdminTravelResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		LoginID             string `json:"loginId"`
		TripID              string `json:"tripId"`
		TrainTypeName       string `json:"trainTypeName"`
		RouteID             string `json:"routeId"`
		StartStationName    string `json:"startStationName"`
		StationsName        string `json:"stationsName"`
		TerminalStationName string `json:"terminalStationName"`
		StartTime           string `json:"startTime"`
		EndTime             string `json:"endTime"`
	} `json:"data"`
}

func (s *SvcImpl) CreateTravel(request *AdminTravelInfo) (*AdminTravelResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/admintravelservice/admintravel", request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminTravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) UpdateTravel(request *AdminTravelInfo) (*AdminTravelResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/admintravelservice/admintravel", request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminTravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) DeleteTravel(tripId string) (*AdminTravelResponse, error) {
	url := fmt.Sprintf("%s/api/v1/admintravelservice/admintravel/%s", s.BaseUrl, tripId)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminTravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) GetAllTravels() ([]AdminTravelInfo, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/admintravelservice/admintravel", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var travels []AdminTravelInfo
	err = json.Unmarshal(body, &travels)
	if err != nil {
		return nil, err
	}

	return travels, nil
}
