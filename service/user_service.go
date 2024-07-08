package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type UserService interface {
	GetAllUsers() (*GetAllUserResponse, error)
	GetUserByUserName(userName string) (*SingleUserResponse, error)
	GetUserByUserId(userId string) (*SingleUserResponse, error)
	RegisterUser(userDto *AdminUserDto) (*SingleUserResponse, error)
	DeleteUser(userId string) (*SingleUserResponse, error)
	UpdateUser(user *AdminUserDto) (*SingleUserResponse, error)
}

type UserDtoUser struct {
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Gender       int    `json:"gender"`
	DocumentType int    `json:"documentType"`
	DocumentNum  string `json:"documentNum"`
	Email        string `json:"email"`
}

type SingleUserResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   UserDtoUser `json:"data"`
}

type AllUserResponseUser struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []UserDtoUser `json:"data"`
}
type GetAllUserResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []UserDtoUser `json:"data"`
}

func (s *SvcImpl) GetAllUsers() (*GetAllUserResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/userservice/users", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetAllUserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) GetUserByUserName(userName string) (*SingleUserResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/userservice/users/%s", userName), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SingleUserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) GetUserByUserId(userId string) (*SingleUserResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/userservice/users/id/%s", userId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SingleUserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) RegisterUser(userDto *AdminUserDto) (*SingleUserResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/userservice/users/register", userDto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SingleUserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) DeleteUser(userId string) (*SingleUserResponse, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/userservice/users/%s", userId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SingleUserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}

func (s *SvcImpl) UpdateUser(user *AdminUserDto) (*SingleUserResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/userservice/users", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result SingleUserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("body: %v", string(body)))
	}
	return &result, nil
}
