package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type AuthService interface {
	ReqUserLogin(input *UserLoginInfoReq) (*UserLoginInfoResp, error)
	ReqUserCreate(input *UserCreateInfoReq) (*UserCreateInfoResp, error)
	ReqUserDelete(userid string) (*UserDeleteInfoResp, error)
}
type UserLoginInfoResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		UserId   string `json:"userId"`
		Username string `json:"username"`
		Token    string `json:"token"`
	} `json:"data"`
}
type UserLoginInfoReq struct {
	Password         string `json:"password"`
	UserName         string `json:"username"`
	VerificationCode string `json:"verificationCode"`
}

type UserCreateInfoReq struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
}
type UserCreateInfoResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		UserId   string `json:"userId"`
		UserName string `json:"userName"`
		Password string `json:"password"`
	} `json:"data"`
}

type UserDeleteInfoResp struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (s *SvcImpl) ReqUserLogin(input *UserLoginInfoReq) (*UserLoginInfoResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/users/login", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result UserLoginInfoResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Data.Token != "" {
		s.cli.AddHeader("Authorization", fmt.Sprintf("Bearer %s", result.Data.Token))
	}
	return &result, nil
}

func (s *SvcImpl) ReqUserCreate(input *UserCreateInfoReq) (*UserCreateInfoResp, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/auth", input)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result UserCreateInfoResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) ReqUserDelete(userid string) (*UserDeleteInfoResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/users/%s", userid), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result UserDeleteInfoResp

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
