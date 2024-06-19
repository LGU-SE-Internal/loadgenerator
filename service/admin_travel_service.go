package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type TravelInfo struct {
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

type TravelResponse struct {
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

func (s *SvcImpl) CreateTravel(request *TravelInfo) (*TravelResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/admintravelservice/admintravel", request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) UpdateTravel(request *TravelInfo) (*TravelResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/admintravelservice/admintravel", request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeleteTravel(tripId string) (*TravelResponse, error) {
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

	var result TravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetAllTravels() ([]TravelInfo, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/admintravelservice/admintravel", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var travels []TravelInfo
	err = json.Unmarshal(body, &travels)
	if err != nil {
		return nil, err
	}

	return travels, nil
}
