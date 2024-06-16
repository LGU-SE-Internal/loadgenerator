package service

import (
	"encoding/json"
	"fmt"
	"io"
)

// Define the request and response structs that will be used in the service methods
type UserDto_user struct {
	UserID       string `json:"userId"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	Gender       int    `json:"gender"`
	DocumentType int    `json:"documentType"`
	DocumentNum  string `json:"documentNum"`
	Email        string `json:"email"`
}

//	type UserResponse struct {
//		Status int    `json:"status"`
//		Msg    string `json:"msg"`
//		Data   []UserDto `json:"data"`
//	}
type SingleUserResponse struct {
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

type UserRegisterResponse struct {
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

//
//type Response struct {
//	Status int    `json:"status"`
//	Msg    string `json:"msg"`
//	Data   interface{} `json:"data"`
//}

type AllUserResponse_user struct {
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

func (s *SvcImpl) GetAllUsers_user() (*UserResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/userservice/users", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result UserResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
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
		return nil, err
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
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) RegisterUser(userDto *UserDto) (*UserRegisterResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/userservice/users/register", userDto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result UserRegisterResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) DeleteUser_user(userId string) (*Response, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/userservice/users/%s", userId), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Response

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SvcImpl) UpdateUser_user(user *UserDto) (*Response, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/userservice/users", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result Response

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
