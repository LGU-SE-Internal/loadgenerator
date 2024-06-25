package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func (s *SvcImpl) ReqPay_InsidePayment(input *TripPayment) (*TripPaymentResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TripPaymentResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqCreateAccount(input *AccountInfo) (*TripPaymentResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/account", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TripPaymentResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqPayDifference(input *TripPayment) (*TripPaymentResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/difference", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TripPaymentResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryAccount() (*TripPaymentArrResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/account", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TripPaymentArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqDrawBack(userId string, money string) (*MoneyResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/drawback/"+userId+"/"+money, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result MoneyResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryAddMoney() (*MoneyResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/money", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result MoneyResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryInsidePayment() (*TripPaymentArrResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/payment", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TripPaymentArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqAddMoney_Inside(userId string, money string) (*TripPaymentResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/inside_pay_service/inside_payment/"+userId+"/"+money, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TripPaymentResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
