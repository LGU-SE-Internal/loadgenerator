package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// PreserveService defines the methods that the service should implement
type PreserveService interface {
	Preserve(orderTicketsInfo *OrderTicketsInfo) (*PreserveResponse, error)
}

// OrderTicketsInfo represents the order tickets info structure
type OrderTicketsInfo struct {
	AccountID       string  `json:"accountId"`
	ContactsID      string  `json:"contactsId"`
	TripID          string  `json:"tripId"`
	SeatType        int     `json:"seatType"`
	LoginToken      string  `json:"loginToken"`
	Date            string  `json:"date"`
	From            string  `json:"from"`
	To              string  `json:"to"`
	Assurance       int     `json:"assurance"`
	FoodType        int     `json:"foodType"`
	StationName     string  `json:"stationName"`
	StoreName       string  `json:"storeName"`
	FoodName        string  `json:"foodName"`
	FoodPrice       float64 `json:"foodPrice"`
	HandleDate      string  `json:"handleDate"`
	ConsigneeName   string  `json:"consigneeName"`
	ConsigneePhone  string  `json:"consigneePhone"`
	ConsigneeWeight float64 `json:"consigneeWeight"`
	IsWithin        bool    `json:"isWithin"`
}

// PreserveResponse represents the response structure for a preserve request
type PreserveResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) Preserve(orderTicketsInfo *OrderTicketsInfo) (*PreserveResponse, error) {
	url := fmt.Sprintf("%s/api/v1/preserveservice/preserve", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, orderTicketsInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result PreserveResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
