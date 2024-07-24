package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// TravelService defines the methods that the service should implement
type TravelService interface {
	GetTrainTypeByTripId(tripId string) (*GetTrainTypeByTripIdResponse, error)
	GetRouteByTripId(tripId string) (*GetRouteByTripIdResponse, error)
	GetTripsByRouteId(routeIds []string) (*GetTripsByRouteIdResponse, error)
	CreateTrip(travelInfo *TravelInfo) (*TripResponse, error)
	RetrieveTravel(tripId string) (*TravelInfo, error)
	UpdateTrip(travelInfo *TravelInfo) (*TripResponse, error)
	DeleteTrip(tripId string) (*TripResponse, error)
	QueryInfo(tripInfo TripInfo) (*QueryInfoResponse, error)
	QueryInfoInParallel(tripInfo TripInfo) (*QueryInfoInParallelTripResponse, error)
	GetTripAllDetailInfo(tripId GetTripDetailReq) (*GetTripAllDetailInfoResponse, error)
	QueryAll() (*QueryAllTravelInfo, error)
	AdminQueryAll() (*AdminQueryAllTravelInfo, error)
}

// TravelInfo represents the travel information
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
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id     string `json:"id"`
		TripId struct {
			Type   string `json:"type"`
			Number string `json:"number"`
		} `json:"tripId"`
		TrainTypeName       string `json:"trainTypeName"`
		RouteId             string `json:"routeId"`
		StartStationName    string `json:"startStationName"`
		StationsName        string `json:"stationsName"`
		TerminalStationName string `json:"terminalStationName"`
		StartTime           string `json:"startTime"`
		EndTime             string `json:"endTime"`
	} `json:"data"`
}

type GetTrainTypeByTripIdResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) GetTrainTypeByTripId(tripId string) (*GetTrainTypeByTripIdResponse, error) {
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

	var result GetTrainTypeByTripIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type GetRouteByTripIdResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id           string      `json:"id"`
		Stations     interface{} `json:"stations"`
		Distances    interface{} `json:"distances"`
		StartStation interface{} `json:"startStation"`
		EndStation   interface{} `json:"endStation"`
	} `json:"data"`
}

func (s *SvcImpl) GetRouteByTripId(tripId string) (*GetRouteByTripIdResponse, error) {
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

	var result GetRouteByTripIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type GetTripsByRouteIdResponse struct {
	Status int             `json:"status"`
	Msg    string          `json:"msg"`
	Data   [][]interface{} `json:"data"`
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

//type QueryInfoResponse struct {
//	Status int `json:"status"`
//	Data   []struct {
//		SeatNo       int    `json:"seatNo"`
//		StartStation string `json:"startStation"`
//		DestStation  string `json:"destStation"`
//	} `json:"data"`
//}

type QueryInfoResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []interface{} `json:"data"`
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type GetTripAllDetailInfoResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Status       bool        `json:"status"`
		Message      interface{} `json:"message"`
		TripResponse interface{} `json:"tripResponse"`
		Trip         interface{} `json:"trip"`
	} `json:"data"`
}
type GetTripDetailReq struct {
	From       string `json:"from"`
	To         string `json:"to"`
	TravelDate string `json:"travelDate"`
	TripId     string `json:"tripId"`
}

func (s *SvcImpl) GetTripAllDetailInfo(trip GetTripDetailReq) (*GetTripAllDetailInfoResponse, error) {
	url := fmt.Sprintf("%s/api/v1/travelservice/trip_detail", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, trip)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetTripAllDetailInfoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

type QueryAllTravelInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id     string `json:"id"`
		TripId struct {
			Type   string `json:"type"`
			Number string `json:"number"`
		} `json:"tripId"`
		TrainTypeName       string `json:"trainTypeName"`
		RouteId             string `json:"routeId"`
		StartStationName    string `json:"startStationName"`
		StationsName        string `json:"stationsName"`
		TerminalStationName string `json:"terminalStationName"`
		StartTime           string `json:"startTime"`
		EndTime             string `json:"endTime"`
	} `json:"data"`
}

type AdminQueryAllTravelInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Trip struct {
			Id     string `json:"id"`
			TripId struct {
				Type   *string `json:"type"`
				Number string  `json:"number"`
			} `json:"tripId"`
			TrainTypeName       string `json:"trainTypeName"`
			RouteId             string `json:"routeId"`
			StartStationName    string `json:"startStationName"`
			StationsName        string `json:"stationsName"`
			TerminalStationName string `json:"terminalStationName"`
			StartTime           string `json:"startTime"`
			EndTime             string `json:"endTime"`
		} `json:"trip"`
		TrainType *struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			EconomyClass int    `json:"economyClass"`
			ConfortClass int    `json:"confortClass"`
			AverageSpeed int    `json:"averageSpeed"`
		} `json:"trainType"`
		Route struct {
			Id           string      `json:"id"`
			Stations     interface{} `json:"stations"`
			Distances    interface{} `json:"distances"`
			StartStation interface{} `json:"startStation"`
			EndStation   interface{} `json:"endStation"`
		} `json:"route"`
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) AdminQueryAll() (*AdminQueryAllTravelInfo, error) {
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

	var result AdminQueryAllTravelInfo
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}
