package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type ConsignService interface {
	InsertConsignRecord(consign *Consign) (*ConsignResponse, error)
	UpdateConsignRecord(consign *Consign) (*ConsignResponse, error)
	QueryByAccountId(accountId string) (*AllConsignResponse, error)
	QueryByOrderId(orderId string) (*QueryByOrderIdResponse, error)
	QueryByConsignee(consignee string) (*QueryByConsigneeResponse, error)
}

// Define the request and response structs that will be used in the service methods
type Consign struct {
	ID         string  `json:"id"`         // id主键改成uuid类型的 自定义生成策略
	OrderID    string  `json:"orderId"`    // 这次托运关联订单
	AccountID  string  `json:"accountId"`  // 这次托运关联的账户
	HandleDate string  `json:"handleDate"` // 处理日期
	TargetDate string  `json:"targetDate"` // 目标日期
	From       string  `json:"from"`       // 出发地
	To         string  `json:"to"`         // 目的地
	Consignee  string  `json:"consignee"`  // 收货人
	Phone      string  `json:"phone"`      // 电话
	Weight     float64 `json:"weight"`     // 重量
	IsWithin   bool    `json:"isWithin"`   // 是否在区域内
}

type ConsignResponse struct {
	Status int     `json:"status"`
	Msg    string  `json:"msg"`
	Data   Consign `json:"data"`
}

type AllConsignResponse struct {
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	Data   []Consign `json:"data"`
}

func (s *SvcImpl) InsertConsignRecord(consign *Consign) (*ConsignResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/consignservice/consigns", consign)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ConsignResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) UpdateConsignRecord(consign *Consign) (*ConsignResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/consignservice/consigns", consign)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ConsignResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SvcImpl) QueryByAccountId(accountId string) (*AllConsignResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/consignservice/consigns/account/%s", accountId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AllConsignResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type QueryByOrderIdResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id         string  `json:"id"`
		OrderId    string  `json:"orderId"`
		AccountId  string  `json:"accountId"`
		HandleDate string  `json:"handleDate"`
		TargetDate string  `json:"targetDate"`
		From       string  `json:"from"`
		To         string  `json:"to"`
		Consignee  string  `json:"consignee"`
		Phone      string  `json:"phone"`
		Weight     float64 `json:"weight"`
		Price      float64 `json:"price"`
	} `json:"data"`
}

func (s *SvcImpl) QueryByOrderId(orderId string) (*QueryByOrderIdResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/consignservice/consigns/order/%s", orderId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryByOrderIdResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type QueryByConsigneeResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id         string  `json:"id"`
		OrderId    string  `json:"orderId"`
		AccountId  string  `json:"accountId"`
		HandleDate string  `json:"handleDate"`
		TargetDate string  `json:"targetDate"`
		From       string  `json:"from"`
		To         string  `json:"to"`
		Consignee  string  `json:"consignee"`
		Phone      string  `json:"phone"`
		Weight     float64 `json:"weight"`
		Price      float64 `json:"price"`
	} `json:"data"`
}

func (s *SvcImpl) QueryByConsignee(consignee string) (*QueryByConsigneeResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/consignservice/consigns/%s", consignee), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result QueryByConsigneeResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
