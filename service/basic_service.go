package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type BasicService interface {
	QueryForTravel(info *Travel) (interface{}, error)
	QueryForTravels(infos []Travel) (interface{}, error)
	QueryForStationId(stationName string) (interface{}, error)
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
	Type   Type   `json:"type"`
	Number string `json:"number"`
}

type Trip struct {
	ID                  string `json:"id"`
	TripID              TripId `json:"tripId"`
	TrainTypeName       string `json:"trainTypeName"`
	RouteID             string `json:"routeId"`
	StartStationName    string `json:"startStationName"`
	StationsName        string `json:"stationsName"`
	TerminalStationName string `json:"terminalStationName"`
	StartTime           string `json:"startTime"`
	EndTime             string `json:"endTime"`
}

type Travel struct {
	Trip          Trip   `json:"trip"`
	StartPlace    string `json:"startPlace"`
	EndPlace      string `json:"endPlace"`
	DepartureTime string `json:"departureTime"`
}

func (s *SvcImpl) QueryForTravel(info *Travel) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) QueryForTravels(infos []Travel) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) QueryForStationId(stationName string) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
