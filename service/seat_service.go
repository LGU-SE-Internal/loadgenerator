package service

import (
	"encoding/json"
	"io"
)

type SeatService interface {
	ReqSeatCreate(input *SeatCreateInfoReq) (*SeatCreateInfoResp, error)
	ReqGetTicketLeft(input *SeatCreateInfoReq) (*TicketLeftResp, error)
}
type SeatCreateInfoReq struct {
	TravelDate  string   `form:"travelDate" json:"travelDate" binding:"required"`
	TrainNumber string   `form:"trainNumber" json:"trainNumber" binding:"required"`
	DestStation string   `form:"destStation" json:"destStation" binding:"required"`
	SeatType    int      `form:"seatType" json:"seatType" binding:"required"`
	TotalNum    int      `form:"totalNum" json:"totalNum" binding:"required"`
	Stations    []string `json:"stations" binding:"required"`
}
type SeatCreateInfoResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		SeatNo       int    `json:"seatNo"`
		StartStation string `json:"startStation"`
		DestStation  string `json:"destStation"`
	} `json:"data"`
}

type TicketLeftResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   int    `json:"data"`
}

func (s *SvcImpl) ReqSeatCreate(input *SeatCreateInfoReq) (*SeatCreateInfoResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/seatservice/seats", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SeatCreateInfoResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetTicketLeft(input *SeatCreateInfoReq) (*TicketLeftResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/seatservice/seats/left_tickets", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TicketLeftResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
