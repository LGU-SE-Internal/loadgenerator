package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type StationService interface {
	QueryStations() (*GetStationResponse, error)
	CreateStation(input *Station) (*StationCreateResponse, error)
	UpdateStation(input *Station) (*StationUpdateResponse, error)
	DeleteStation(stationId string) (*DeleteStationResponse, error)
	QueryStationIdByName(stationName string) (*StationQueryIdByNameResponse, error)
	QueryStationIdsByNames(stationNameList []string) (*QueryStationIdsByNamesResponse, error)
	QueryStationNameById(stationId string) (*QueryStationNameByIdResponse, error)
	QueryStationNamesByIds(stationIdList []string) (*QueryStationNamesByIdsResponse, error)
}
type Station struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	StayTime int    `json:"stayTime"`
}

type DeleteStationResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		StayTime int    `json:"stayTime"`
	} `json:"data"`
}
type GetStationResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		StayTime int    `json:"stayTime"`
	} `json:"data"`
}

type StationCreateResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		StayTime int    `json:"stayTime"`
	} `json:"data"`
}

type StationUpdateResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		StayTime int    `json:"stayTime"`
	} `json:"data"`
}

type StationQueryIdByNameResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type QueryStationIdsByNamesResponse struct {
	Status int               `json:"status"`
	Msg    string            `json:"msg"`
	Data   map[string]string `json:"data"`
}

type QueryStationNameByIdResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type QueryStationNamesByIdsResponse struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   []string `json:"data"`
}

func (s *SvcImpl) QueryStations() (*GetStationResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/stationservice/stations", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetStationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) CreateStation(input *Station) (*StationCreateResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/stationservice/stations", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result StationCreateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) UpdateStation(input *Station) (*StationUpdateResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/stationservice/stations", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result StationUpdateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) DeleteStation(stationId string) (*DeleteStationResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/stationservice/stations/%s", stationId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DeleteStationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryStationIdByName(stationName string) (*StationQueryIdByNameResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/stationservice/stations/id/%s", stationName), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result StationQueryIdByNameResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryStationIdsByNames(stationNameList []string) (*QueryStationIdsByNamesResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/stationservice/stations/idlist", stationNameList)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryStationIdsByNamesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryStationNameById(stationId string) (*QueryStationNameByIdResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/stationservice/stations/name/%s", stationId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryStationNameByIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryStationNamesByIds(stationIdList []string) (*QueryStationNamesByIdsResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/stationservice/stations/namelist", stationIdList)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryStationNamesByIdsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
