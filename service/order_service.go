package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type Seat struct {
	TravelDate   string   `json:"travelDate" validate:"required"`
	TrainNumber  string   `json:"trainNumber" validate:"required"`
	StartStation string   `json:"startStation" validate:"required"`
	DestStation  string   `json:"destStation" validate:"required"`
	SeatType     int      `json:"seatType" validate:"required"`
	TotalNum     int      `json:"totalNum"`
	Stations     []string `json:"stations"`
}

type Ticket struct {
	status string `json:"travelDate" validate:"required"`
	msg    string `json:"trainNumber" validate:"required"`
	Data   string `json:"startStation" validate:"required"`
}

type Order struct {
	status string `json:"travelDate" validate:"required"`
	msg    string `json:"trainNumber" validate:"required"`
	Data   string `json:"startStation" validate:"required"`
}

type OrderInfo struct {
	status string `json:"travelDate" validate:"required"`
	msg    string `json:"trainNumber" validate:"required"`
	Data   string `json:"startStation" validate:"required"`
}

func (s *SvcImpl) GetTicketListByDateAndTripId(seatRequest Seat) ([]Ticket, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/tickets", seatRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []Ticket
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SvcImpl) CreateNewOrder(order Order) (*Order, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order", order)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) AddNewOrder(order Order) (*Order, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/admin", order)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) QueryOrders(orderInfo OrderInfo) ([]Order, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/query", orderInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SvcImpl) QueryOrdersForRefresh(orderInfo OrderInfo) ([]Order, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/orderservice/order/refresh", orderInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *SvcImpl) CalculateSoldTicket(travelDate, trainNumber string) (int, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/orderservice/order/%s/%s", s.BaseUrl, travelDate, trainNumber), nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var result int
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *SvcImpl) GetOrderPrice(orderId string) (float64, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/orderservice/order/price/%s", s.BaseUrl, orderId), nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var result float64
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *SvcImpl) PayOrder(orderId string) (*Order, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/orderservice/order/orderPay/%s", s.BaseUrl, orderId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) GetOrderById(orderId string) (*Order, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/orderservice/order/%s", s.BaseUrl, orderId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ModifyOrder(orderId string, status int) (*Order, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/orderservice/order/status/%s/%d", s.BaseUrl, orderId, status), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) SecurityInfoCheck(checkDate, accountId string) (bool, error) {
	resp, err := s.cli.SendRequest("GET", fmt.Sprintf("%s/api/v1/orderservice/order/security/%s/%s", s.BaseUrl, checkDate, accountId), nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var result bool
	err = json.Unmarshal(body, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (s *SvcImpl) SaveOrderInfo(orderInfo Order) (*Order, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderservice/order", orderInfo)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) UpdateOrder(order Order) (*Order, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/orderservice/order/admin", order)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) DeleteOrder(orderId string) (*Order, error) {
	resp, err := s.cli.SendRequest("DELETE", fmt.Sprintf("%s/api/v1/orderservice/order/%s", s.BaseUrl, orderId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) GetAllOrders() ([]Order, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/orderservice/order", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result []Order
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
