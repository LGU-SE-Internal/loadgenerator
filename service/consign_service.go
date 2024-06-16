package service

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//)
//
//type Consign struct {
//	ID string `json:"id"`
//	// other fields...
//}
//
//func (s *SvcImpl) InsertConsignRecord(request Consign) (interface{}, error) {
//	url := fmt.Sprintf("%s/api/v1/consignservice/consigns", s.BaseUrl)
//	resp, err := s.sendRequest("POST", url, request)
//	if err != nil {
//		return nil, err
//	}
//	return s.parseResponse(resp)
//}
//
//func (s *SvcImpl) UpdateConsignRecord(request Consign) (interface{}, error) {
//	url := fmt.Sprintf("%s/api/v1/consignservice/consigns", s.BaseUrl)
//	resp, err := s.sendRequest("PUT", url, request)
//	if err != nil {
//		return nil, err
//	}
//	return s.parseResponse(resp)
//}
//
//func (s *SvcImpl) QueryByAccountId(id string) (interface{}, error) {
//	url := fmt.Sprintf("%s/api/v1/consignservice/consigns/account/%s", s.BaseUrl, id)
//	resp, err := s.sendRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//	return s.parseResponse(resp)
//}
//
//func (s *SvcImpl) QueryByOrderId(id string) (interface{}, error) {
//	url := fmt.Sprintf("%s/api/v1/consignservice/consigns/order/%s", s.BaseUrl, id)
//	resp, err := s.sendRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//	return s.parseResponse(resp)
//}
//
//func (s *SvcImpl) QueryByConsignee(consignee string) (interface{}, error) {
//	url := fmt.Sprintf("%s/api/v1/consignservice/consigns/%s", s.BaseUrl, consignee)
//	resp, err := s.sendRequest("GET", url, nil)
//	if err != nil {
//		return nil, err
//	}
//	return s.parseResponse(resp)
//}
//
//func (s *SvcImpl) sendRequest(method, url string, data interface{}) (*http.RouteResponse, error) {
//	var body io.Reader
//	if data != nil {
//		jsonData, err := json.Marshal(data)
//		if err != nil {
//			return nil, err
//		}
//		body = bytes.NewReader(jsonData)
//	}
//
//	req, err := http.NewRequest(method, url, body)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//	return s.cli.Do(req)
//}
//
//func (s *SvcImpl) parseResponse(resp *http.RouteResponse) (interface{}, error) {
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var result interface{}
//	err = json.Unmarshal(body, &result)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
