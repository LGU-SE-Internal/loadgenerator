package service

import (
	"encoding/json"
	"io"
)

// Define the structs for the request and response
type RoutePlanInfo struct {
	StartStation string `json:"startStation"`
	EndStation   string `json:"endStation"`
	Num          int    `json:"num"`
	TravelDate   string `json:"travelDate"`
}

type RoutePlanResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Routes Route `json:"routes,omitempty"`
	} `json:"data,omitempty"`
}

type Route struct {
	// Define fields based on the expected response
}

func (s *SvcImpl) GetCheapestRoutes(input *RoutePlanInfo) (*RoutePlanResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/routeplanservice/routePlan/cheapestRoute", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result RoutePlanResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) GetQuickestRoutes(input *RoutePlanInfo) (*RoutePlanResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/routeplanservice/routePlan/quickestRoute", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result RoutePlanResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) GetMinStopStations(input *RoutePlanInfo) (*RoutePlanResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/routeplanservice/routePlan/minStopStations", input)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result RoutePlanResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
