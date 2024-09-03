package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type CancelService interface {
	ReqCalculate(orderId string) (*ReqCalculateResp, error)
	ReqCancelTicket(orderId string, loginId string) (*CancelTicketResp, error)
}

type ReqCalculateResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type CancelTicketResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) ReqCalculate(orderId string) (*ReqCalculateResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/cancelservice/cancel/refound/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ReqCalculateResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqCancelTicket(orderId string, loginId string) (*CancelTicketResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/cancelservice/cancel/"+orderId+"/"+loginId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result CancelTicketResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
