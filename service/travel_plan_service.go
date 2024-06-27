package service

import (
	"encoding/json"
	"io"
)

func (s *SvcImpl) ReqGetByCheapest(input *TravelQueryInfo) (*TravelQueryArrResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travelplanservice/travelPlan/cheapest", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TravelQueryArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetByMinStation(input *TravelQueryInfo) (*TravelQueryArrResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travelplanservice/travelPlan/minStation", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TravelQueryArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqGetByQuickest(input *TravelQueryInfo) (*TravelQueryArrResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travelplanservice/travelPlan/quickest", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TravelQueryArrResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqTransferResult(input *TransferTravelQueryInfo) (*TravelQueryResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/travelplanservice/travelPlan/transferResult", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result TravelQueryResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
