package service

import (
	"encoding/json"
	"fmt"
	"io"
)

func (s *SvcImpl) VerifyCode(verifyCode string) (bool, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/verifycode/verify/%s", verifyCode), nil)
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
