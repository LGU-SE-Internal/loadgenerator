package service

import (
	"encoding/json"
	"io"
	"strconv"
)

func (s *SvcImpl) ReqFindAllOeder_Other() (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther", nil)
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

func (s *SvcImpl) ReqCreateNewOeder_Other(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqSaveOrderInfo_Other(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderOtherService/orderOther", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqAddCreateNewOrder_Other(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/admin", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateOrder_OrderService_Other(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderOtherService/orderOther/admin", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqPayOrder_Other(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/orderpay/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetOrderPrice_Other(orderId string) (*GetOrderPriceResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/price/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetOrderPriceResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryOrders_Other(input *Qi) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/query", input)
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

func (s *SvcImpl) ReqQueryOrderForRefresh_Other(input *Qi) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/refresh", input)
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

func (s *SvcImpl) ReqSecurityInfoCheck_Other(checkDate string, accountId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/security/"+checkDate+"/"+accountId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqModifyOrder_Other(orderId string, status int) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/status/"+orderId+"/"+strconv.Itoa(status), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetTicketsList_Other(input *Seat) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderOtherService/orderOther/tickets", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqDeleteOrder_OrderService_Other(orderId string) (*DataStringResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+"/api/v1/orderOtherService/orderOther/"+orderId, nil)
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

func (s *SvcImpl) ReqGetOrderById_Other(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/"+orderId, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqCalculateSoldTicket_Other(travelDate string, travelNumber string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderOtherService/orderOther/"+travelDate+"/"+travelNumber, nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result OrderResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
