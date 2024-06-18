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
//}

// Route represents the route structure
//type Route struct {
//	ID       string `json:"id"`
//	Name     string `json:"name"`
//	Stations []string `json:"stations"`
//}

// TravelInfo represents the travel information
type TravelInfo struct {
	TripId string `json:"tripId"`
	// Add other fields as necessary
}

// TripInfo represents the trip information
type TripInfo struct {
	StartPlace    string `json:"startPlace"`
	EndPlace      string `json:"endPlace"`
	DepartureTime string `json:"departureTime"`
	// Add other fields as necessary
}

// TripAllDetailInfo represents the trip all detail information
type TripAllDetailInfo struct {
	TripId string `json:"tripId"`
	// Add other fields as necessary
}

// TripResponse represents the trip response
type TripResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// TravelService defines the methods that the service should implement
type TravelService interface {
	GetTrainTypeByTripId(tripId string) (*TrainType, error)
	GetRouteByTripId(tripId string) (*Route, error)
	GetTripsByRouteId(routeIds []string) (*GetTripsByRouteIdResponse, error)
	CreateTrip(travelInfo *TravelInfo) (*TripResponse, error)
	RetrieveTravel(tripId string) (*TravelInfo, error)
	UpdateTrip(travelInfo *TravelInfo) (*TripResponse, error)
	DeleteTrip(tripId string) (*TripResponse, error)
	QueryInfo(tripInfo TripInfo) (*QueryInfoResponse, error)
	QueryInfoInParallel(tripInfo TripInfo) (*QueryInfoInParallelTripResponse, error)
	GetTripAllDetailInfo(tripAllDetailInfo TripAllDetailInfo) (*TripResponse, error)
	QueryAll() (*QueryAllTravelInfo, error)
	AdminQueryAll() (*QueryAllTravelInfo, error)
}

func (s *SvcImpl) GetTrainTypeByTripId(tripId string) (*TrainType, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/train_types/%s", s.BaseUrl, tripId)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TrainType
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetRouteByTripId(tripId string) (*Route, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/routes/%s", s.BaseUrl, tripId)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Route
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type GetTripsByRouteIdResponse struct {
	Status int `json:"status"`
	Data   []struct {
		SeatNo       int    `json:"seatNo"`
		StartStation string `json:"startStation"`
		DestStation  string `json:"destStation"`
	} `json:"data"`
}

func (s *SvcImpl) GetTripsByRouteId(routeIds []string) (*GetTripsByRouteIdResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips/routes", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, routeIds)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetTripsByRouteIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) CreateTrip(travelInfo *TravelInfo) (*TripResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, travelInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) RetrieveTravel(tripId string) (*TravelInfo, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips/%s", s.BaseUrl, tripId)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TravelInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) UpdateTrip(travelInfo *TravelInfo) (*TripResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips", s.BaseUrl)
	resp, err := s.cli.SendRequest("PUT", url, travelInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) DeleteTrip(tripId string) (*TripResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips/%s", s.BaseUrl, tripId)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type QueryInfoResponse struct {
	Status int `json:"status"`
	Data   []struct {
		SeatNo       int    `json:"seatNo"`
		StartStation string `json:"startStation"`
		DestStation  string `json:"destStation"`
	} `json:"data"`
}

func (s *SvcImpl) QueryInfo(tripInfo TripInfo) (*QueryInfoResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips/left", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, tripInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryInfoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type QueryInfoInParallelTripResponse struct {
	Status int `json:"status"`
	Data   []struct {
		SeatNo       int    `json:"seatNo"`
		StartStation string `json:"startStation"`
		DestStation  string `json:"destStation"`
	} `json:"data"`
}

func (s *SvcImpl) QueryInfoInParallel(tripInfo TripInfo) (*QueryInfoInParallelTripResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips/left_parallel", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, tripInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryInfoInParallelTripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) GetTripAllDetailInfo(tripAllDetailInfo TripAllDetailInfo) (*TripResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trip_detail", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, tripAllDetailInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result TripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type QueryAllTravelInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		SeatNo       int    `json:"seatNo"`
		StartStation string `json:"startStation"`
		DestStation  string `json:"destStation"`
	} `json:"data"`
}

func (s *SvcImpl) QueryAll() (*QueryAllTravelInfo, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trips", s.BaseUrl)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryAllTravelInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) AdminQueryAll() (*QueryAllTravelInfo, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/admin_trip", s.BaseUrl)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryAllTravelInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
