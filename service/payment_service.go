package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqPay(input *Payment) (*PaymentResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/paymentservice/payment", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PaymentResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqAddMoney(input *Payment) (*PaymentResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/paymentservice/payment/money", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PaymentResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryPayment() (*PaymentArrResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/paymentservice/payment", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result PaymentArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
