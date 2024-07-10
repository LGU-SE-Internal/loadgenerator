package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqOrderCancelSuccess(input *TicketOrder) (*bool, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/notifyservice/notification/order_cancel_success", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result bool

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqOrderChangedSuccess(input *TicketOrder) (*bool, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/notifyservice/notification/order_changed_success", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result bool

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqOrderCreateSuccess(input *TicketOrder) (*bool, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/notifyservice/notification/order_create_success", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result bool

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqPreserveSuccess(input *TicketOrder) (*bool, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/notifyservice/notification/preserve_success", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result bool

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqTestSendMail() (*bool, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/notifyservice/test_send_mail", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result bool

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqTestSend() (*bool, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/notifyservice/test_send_mq", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result bool

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
