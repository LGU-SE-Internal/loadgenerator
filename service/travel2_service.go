package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Travel2Service interface {
	GetTrain2TypeByTripId(tripId string) (*GetTrainTypeByTripId2Response, error)
	GetRouteByTrip2Id(tripId string) (*GetRouteByTripIdResponse, error)
	GetTrip2ByRoute(routeIds []string) (*GetTripByRouteIdResponse, error)
	CreateTrip2(travelInfo *TravelInfo) (*CreateTripResponse, error)
	RetrieveTrip2(tripId string) (*RetrieveTripResponse, error)
	UpdateTrip2(travelInfo *TravelInfo) (*UpdateTripResponse, error)
	DeleteTrip2(tripId string) (*DeleteTripResponse, error)
	QueryByBatch(tripInfo *TripInfo) (*QueryByBatchResponse, error)
	GetTrip2AllDetailInfo(tripAllDetailInfo *Trip2AllDetailInfo) (*GetTripAllDetailInfoResponse, error)
	QueryAllTravel() (*QueryAllResponse, error)
	AdminQueryAllTravel() (*AdminQueryAllResponse, error)
}

type Travel2Info struct {
	TripId string `json:"tripId"`
	// Add other fields as needed
}

type Trip2Info struct {
	StartPlace    string `json:"startPlace"`
	EndPlace      string `json:"endPlace"`
	DepartureTime string `json:"departureTime"`
	// Add other fields as needed
}

type Trip2AllDetailInfo struct {
	TripID     string `json:"tripId"`
	TravelDate string `json:"travelDate"`
	From       string `json:"from"`
	To         string `json:"to"`
}

//type GetTrainTypeByTripId2Response struct {
//	Status int    `json:"status"`
//	Msg    string `json:"msg"`
//	Data   string `json:"data"`
//}

type GetTrainTypeByTripId2Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type GetRouteByTrip2IdResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type GetTripByRouteIdResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type CreateTripResponse struct {
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

type RetrieveTripResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type UpdateTripResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type DeleteTripResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type QueryByBatchResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []interface{} `json:"data"`
}

type GetTrip2AllDetailInfoResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type QueryAllResponse struct {
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

type AdminQueryAllResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Trip struct {
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
		} `json:"trip"`
		TrainType *struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			EconomyClass int    `json:"economyClass"`
			ConfortClass int    `json:"confortClass"`
			AverageSpeed int    `json:"averageSpeed"`
		} `json:"trainType"`
		Route *struct {
			Id           string   `json:"id"`
			Stations     []string `json:"stations"`
			Distances    []int    `json:"distances"`
			StartStation string   `json:"startStation"`
			EndStation   string   `json:"endStation"`
		} `json:"route"`
	} `json:"data"`
}

func (s *SvcImpl) GetTrain2TypeByTripId(tripId string) (*GetTrainTypeByTripId2Response, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/travel2service/train_types/%s", tripId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetTrainTypeByTripId2Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) GetRouteByTrip2Id(tripId string) (*GetRouteByTripIdResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/travel2service/routes/%s", tripId), nil)
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

func (s *SvcImpl) GetTrip2ByRoute(routeIds []string) (*GetTripByRouteIdResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travel2service/trips/routes", routeIds)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GetTripByRouteIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) CreateTrip2(travelInfo *TravelInfo) (*CreateTripResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travel2service/trips", travelInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CreateTripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) RetrieveTrip2(tripId string) (*RetrieveTripResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/travel2service/trips/%s", tripId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result RetrieveTripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) UpdateTrip2(travelInfo *TravelInfo) (*UpdateTripResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/travel2service/trips", travelInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result UpdateTripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) DeleteTrip2(tripId string) (*DeleteTripResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/travel2service/trips/%s", tripId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result DeleteTripResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) QueryByBatch(tripInfo *TripInfo) (*QueryByBatchResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travel2service/trips/left", tripInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryByBatchResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) GetTrip2AllDetailInfo(tripAllDetailInfo *Trip2AllDetailInfo) (*GetTripAllDetailInfoResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travel2service/trip_detail", tripAllDetailInfo)
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

func (s *SvcImpl) QueryAllTravel() (*QueryAllResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/travel2service/trips", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result QueryAllResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) AdminQueryAllTravel() (*AdminQueryAllResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/travel2service/admin_trip", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminQueryAllResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
