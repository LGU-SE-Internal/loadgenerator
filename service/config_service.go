package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type Config_config struct {
	Name        string `json:"name" validate:"required"`  // @Valid, @Id, @NotNull
	Value       string `json:"value" validate:"required"` // @Valid, @NotNull
	Description string `json:"description" validate:""`   // @Valid
}

type ConfigQueryAllConfigsResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"data"`
}

func (s *SvcImpl) QueryAllConfigs() (*ConfigQueryAllConfigsResponse, error) {
	url := fmt.Sprintf("%s/api/v1/configservice/configs", s.BaseUrl)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ConfigQueryAllConfigsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type CreateConfigResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"data"`
}

func (s *SvcImpl) CreateConfig(info *Config_config) (*CreateConfigResponse, error) {
	url := fmt.Sprintf("%s/api/v1/configservice/configs", s.BaseUrl)
	resp, err := s.cli.SendRequest("POST", url, info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CreateConfigResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type UpdateConfigResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"data"`
}

func (s *SvcImpl) UpdateConfig(info Config_config) (*UpdateConfigResponse, error) {
	url := fmt.Sprintf("%s/api/v1/configservice/configs", s.BaseUrl)
	resp, err := s.cli.SendRequest("PUT", url, info)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result UpdateConfigResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type DeleteConfig_config_serviceResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"data"`
}

func (s *SvcImpl) DeleteConfig_config_service(configName string) (*DeleteConfig_config_serviceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/configservice/configs/%s", s.BaseUrl, configName)
	resp, err := s.cli.SendRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result DeleteConfig_config_serviceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type RetrieveConfigResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`
	} `json:"data"`
}

func (s *SvcImpl) RetrieveConfig(configName string) (*RetrieveConfigResponse, error) {
	url := fmt.Sprintf("%s/api/v1/configservice/configs/%s", s.BaseUrl, configName)
	resp, err := s.cli.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result RetrieveConfigResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
