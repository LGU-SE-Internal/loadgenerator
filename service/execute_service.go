package service

import (
	"encoding/json"
	"io"
)

type ExecuteService interface {
	ReqExecuteTicket(orderId string) (*DataStringResp, error)
	ReqCollectTicket(orderId string) (*DataStringResp, error)
}

func (s *SvcImpl) ReqExecuteTicket(orderId string) (*DataStringResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/executeservice/execute/execute/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DataStringResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqCollectTicket(orderId string) (*DataStringResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/executeservice/execute/collected/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DataStringResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
