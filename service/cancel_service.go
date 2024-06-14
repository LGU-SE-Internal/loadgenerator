package service

import (
	"encoding/json"
	"fmt"
	"io"
)

func (s *SvcImpl) CalculateRefund(orderID string) (interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/cancelservice/cancel/refound/%s", s.BaseUrl, orderID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) CancelOrder(orderID string, loginID string) (interface{}, error) {
	url := fmt.Sprintf("%s/api/v1/cancelservice/cancel/%s/%s", s.BaseUrl, orderID, loginID)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
