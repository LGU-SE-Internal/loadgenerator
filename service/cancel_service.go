package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqCalculate(orderId string) (*DataStringResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/cancelservice/cancel/refound/"+orderId, nil)
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

func (s *SvcImpl) ReqCancelTicket(orderId string, loginId string) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/cancelservice/cancel/"+orderId+"/"+loginId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderArrResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
