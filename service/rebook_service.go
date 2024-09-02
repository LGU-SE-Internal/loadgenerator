package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type ReBookService interface {
	PayDifference(info *RebookInfo) (*RebookResponse, error)
	Rebook(info *RebookInfo) (*RebookResponse, error)
}

type RebookInfo struct {
	LoginID   string `json:"loginId"`
	OrderID   string `json:"orderId"`
	OldTripID string `json:"oldTripId"`
	TripID    string `json:"tripId"`
	SeatType  int    `json:"seatType"`
	Date      string `json:"date"`
}

type RebookResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) PayDifference(info *RebookInfo) (*RebookResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/rebookservice/rebook/difference", info)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result RebookResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) Rebook(info *RebookInfo) (*RebookResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/rebookservice/rebook", info)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result RebookResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
