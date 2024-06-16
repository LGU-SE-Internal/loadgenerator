package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type Station_station struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	StayTime int    `json:"stayTime"`
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		StayTime int    `json:"stayTime"`
	} `json:"data"`
}

type createResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		StayTime int    `json:"stayTime"`
	} `json:"data"`
}

type updateResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type queryIdByNameResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type queryStationIdsByNamesResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Suzhou       string `json:"suzhou"`
		Shijiazhuang string `json:"shijiazhuang"`
	} `json:"data"`
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

func (s *SvcImpl) QueryStations() (*Response, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/stationservice/stations", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) CreateStation(input *Station_station) (*createResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/stationservice/stations", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result createResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) UpdateStation(input *Station_station) (*updateResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/stationservice/stations", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result updateResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) DeleteStation_station(stationId string) (*Response, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/stationservice/stations/%s", stationId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) QueryStationIdByName(stationName string) (*queryIdByNameResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/stationservice/stations/id/%s", stationName), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result queryIdByNameResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) QueryStationIdsByNames(stationNameList []string) (*queryStationIdsByNamesResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/stationservice/stations/idlist", stationNameList)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result queryStationIdsByNamesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}
	return &result, nil
}
