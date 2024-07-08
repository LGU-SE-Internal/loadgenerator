package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type AdminUserService interface {
	AdminAddUser(user *AdminUserDto) (*AdminUserResponse, error)
	AdminUpdateUser(user *AdminUserDto) (*AdminUserResponse, error)
	AdminDeleteUser(userId string) (*AdminDeleteResponseUser, error)
	AdminGetAllUsers() (*AllUserResponseUser, error)
}
type AdminUserDto struct {
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Gender       int    `json:"gender"`
	DocumentType int    `json:"documentType"`
	DocumentNum  string `json:"documentNum"`
	Email        string `json:"email"`
}

type AdminDeleteResponseUser struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type AdminUserResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		UserId       string `json:"userId"`
		UserName     string `json:"userName"`
		Password     string `json:"password"`
		Gender       int    `json:"gender"`
		DocumentType int    `json:"documentType"`
		DocumentNum  string `json:"documentNum"`
		Email        string `json:"email"`
	} `json:"data"`
}

type AdminAllUserResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		UserId       string `json:"userId"`
		UserName     string `json:"userName"`
		Password     string `json:"password"`
		Gender       int    `json:"gender"`
		DocumentType int    `json:"documentType"`
		DocumentNum  string `json:"documentNum"`
		Email        string `json:"email"`
	} `json:"data"`
}

func (s *SvcImpl) AdminAddUser(user *AdminUserDto) (*AdminUserResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/adminuserservice/users", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminUserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) AdminUpdateUser(user *AdminUserDto) (*AdminUserResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/adminuserservice/users", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminUserResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) AdminDeleteUser(userId string) (*AdminDeleteResponseUser, error) {
	url := fmt.Sprintf("%s/api/v1/adminuserservice/users/%s", s.BaseUrl, userId)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AdminDeleteResponseUser
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}

	return &result, nil
}

func (s *SvcImpl) AdminGetAllUsers() (*AllUserResponseUser, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/adminuserservice/users", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users *AllUserResponseUser
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
