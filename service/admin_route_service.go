package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type RouteInfo struct {
	LoginID      string `json:"loginId"`
	StartStation string `json:"startStation"`
	EndStation   string `json:"endStation"`
	StationList  string `json:"stationList"`
	DistanceList string `json:"distanceList"`
	ID           string `json:"id"`
}

type AddResponse struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   []string `json:"data"`
}

//type RouteInfoResp struct {
//	Status int    `json:"status"`
//	Msg    string `json:"msg"`
//	Data   []struct {
//		LoginID      string `json:"loginId"`
//		StartStation string `json:"startStation"`
//		EndStation   string `json:"endStation"`
//		StationList  string `json:"stationList"`
//		DistanceList string `json:"distanceList"`
//		ID           string `json:"id"`
//	} `json:"data"`
//}

type RouteInfoResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id           string   `json:"id"`
		Stations     []string `json:"stations"`
		Distances    []int    `json:"distances"`
		StartStation string   `json:"startStation"`
		EndStation   string   `json:"endStation"`
	} `json:"data"`
}

type RouteDeleteInfoResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) ReqGetAllRoutes() (*RouteInfoResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminrouteservice/adminroute", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result RouteInfoResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

type AdminAddResponse struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   []string `json:"data"`
}

func (s *SvcImpl) ReqAddRoute(input *RouteInfo) (*AdminAddResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminrouteservice/adminroute", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminAddResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (s *SvcImpl) ReqDeleteRoute(routeId string) (*RouteDeleteInfoResp, error) {
	resp, err := s.cli.SendRequest("DELETE", fmt.Sprintf("%s/api/v1/adminrouteservice/adminroute/%s", s.BaseUrl, routeId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result RouteDeleteInfoResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
