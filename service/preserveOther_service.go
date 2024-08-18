package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// PreserveService defines the methods that the service should implement
type PreserveOtherService interface {
	PreserveOther(orderTicketsInfo *OrderTicketsInfo) (*PreserveOtherResponse, error)
}

type PreserveOtherResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) PreserveOther(orderTicketsInfo *OrderTicketsInfo) (*PreserveOtherResponse, error) {
	url := fmt.Sprintf("%s/api/v1/preserveotherservice/preserveOther", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, orderTicketsInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result PreserveOtherResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
