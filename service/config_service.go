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

func (s *SvcImpl) QueryAllConfigs() (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) CreateConfig(info Config_config) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) UpdateConfig(info Config_config) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) DeleteConfig_config_service(configName string) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SvcImpl) RetrieveConfig(configName string) (interface{}, error) {
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

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
