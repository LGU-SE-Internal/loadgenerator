package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type OrderService interface {
	ReqFindAllOrder() (*OrderArrResp, error)
	ReqCreateNewOrder(input *Order) (*OrderResp, error)
	ReqSaveOrderInfo(input *Order) (*OrderResp, error)
	ReqAddCreateNewOrder(input *Order) (*OrderResp, error)
	ReqUpdateOrder_OrderService(input *Order) (*OrderResp, error)
	ReqPayOrder(orderId string) (*OrderResp, error)
	ReqGetOrderPrice(orderId string) (*GetOrderPriceResp, error)
	ReqQueryOrders(input *Qi) (*OrderArrResp, error)
	ReqQueryOrderForRefresh(input *Qi) (*OrderArrResp, error)
	ReqSecurityInfoCheck(checkDate string, accountId string) (*OrderResp, error)
	ReqModifyOrder(orderId string, status int) (*OrderResp, error)
	ReqGetTicketsList(input *Seat) (*OrderResp, error)
	ReqDeleteOrder_OrderService(orderId string) (*OrderResp, error)
	ReqGetOrderById(orderId string) (*OrderResp, error)
	ReqCalculateSoldTicket(travelDate string, travelNumber string) (*OrderResp, error)
}

func (s *SvcImpl) ReqFindAllOrder() (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order", nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqCreateNewOrder(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqSaveOrderInfo(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderservice/order", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqAddCreateNewOrder(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/admin", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqUpdateOrder_OrderService(input *Order) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderservice/order/admin", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqPayOrder(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order/orderpay/"+orderId, nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetOrderPrice(orderId string) (*GetOrderPriceResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order/price/"+orderId, nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryOrders(input *Qi) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/query", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqQueryOrderForRefresh(input *Qi) (*OrderArrResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/refresh", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqSecurityInfoCheck(checkDate string, accountId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order/security/"+checkDate+"/"+accountId, nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqModifyOrder(orderId string, status int) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order/status/"+orderId+"/"+strconv.Itoa(status), nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetTicketsList(input *Seat) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/tickets", input)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqDeleteOrder_OrderService(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+"/api/v1/orderservice/order/"+orderId, nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetOrderById(orderId string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order/"+orderId, nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) ReqCalculateSoldTicket(travelDate string, travelNumber string) (*OrderResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order/"+travelDate+"/"+travelNumber, nil)
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
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
