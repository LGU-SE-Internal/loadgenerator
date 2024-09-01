package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type ReBookService interface {
	PayDifference(info *RebookInfo) (*RoutePlanResponse, error)
	Rebook(info *RebookInfo) (*RoutePlanResponse, error)
}

type RebookInfo struct {
	LoginID   string
	OrderID   string
	OldTripID string
	TripID    string
	SeatType  int
	Date      string
}

func (s *SvcImpl) PayDifference(info *RebookInfo) (*RoutePlanResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/rebookservice/rebook/difference", info)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) Rebook(info *RebookInfo) (*RoutePlanResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/rebookservice/rebook", info)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
