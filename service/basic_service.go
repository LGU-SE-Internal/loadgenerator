package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type BasicService interface {
	QueryForTravel(info *Travel) (*QueryForTravelResponse, error)
	QueryForTravels(infos []*Travel) (*QueryForTravelsResponse, error)
	QueryForStationId(stationName string) (*QueryForStationIdResponse, error)
	QueryTrainService() (*TrainResponseType, error)
}
type Type struct {
	//G("G", 1),
	///**
	// * D
	// */
	//D("D", 2),
	///**
	// * Z
	// */
	//Z("Z",3),
	///**
	// * T
	// */
	//T("T", 4),
	///**
	// * K
	// */
	//K("K", 5);
	Name  string `json:"name"`
	Index int    `json:"index"`
}

type TripId struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Trip struct {
	EndTime             string `json:"endTime"`
	Id                  string `json:"id"`
	RouteId             string `json:"routeId"`
	StartStationName    string `json:"startStationName"`
	StartTime           string `json:"startTime"`
	StationsName        string `json:"stationsName"`
	TerminalStationName string `json:"terminalStationName"`
	TrainTypeName       string `json:"trainTypeName"`
	TripId              TripId `json:"tripId"`
}

type Travel struct {
	DepartureTime string `json:"departureTime"`
	EndPlace      string `json:"endPlace"`
	StartPlace    string `json:"startPlace"`
	Trip          Trip   `json:"trip"`
}

type QueryForTravelResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Status    bool    `json:"status"`
		Percent   float64 `json:"percent"`
		TrainType struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			EconomyClass int    `json:"economyClass"`
			ConfortClass int    `json:"confortClass"`
			AverageSpeed int    `json:"averageSpeed"`
		} `json:"trainType"`
		Route struct {
			Id           string   `json:"id"`
			Stations     []string `json:"stations"`
			Distances    []int    `json:"distances"`
			StartStation string   `json:"startStation"`
			EndStation   string   `json:"endStation"`
		} `json:"route"`
		Prices struct {
			ConfortClass string `json:"confortClass"`
			EconomyClass string `json:"economyClass"`
		} `json:"prices"`
	} `json:"data"`
}

func (s *SvcImpl) QueryForTravel(info *Travel) (*QueryForTravelResponse, error) {
	url := fmt.Sprintf("%s/api/v1/basicservice/basic/travel", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryForTravelResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) QueryTrainService() (*TrainResponseType, error) {
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

type QueryForTravelsResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) QueryForTravels(infos []*Travel) (*QueryForTravelsResponse, error) {
	url := fmt.Sprintf("%s/api/v1/basicservice/basic/travels", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, infos)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryForTravelsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type QueryForStationIdResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) QueryForStationId(stationName string) (*QueryForStationIdResponse, error) {
	url := fmt.Sprintf("%s/api/v1/basicservice/basic/%s", s.BaseUrl, stationName)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryForStationIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}
