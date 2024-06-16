package service

import (
	"testing"
)

func TestVerifyCodeService_VerifyCode(t *testing.T) {
	cli, _ := GetAdminClient()
	verifyCode := "test-code"
	result, err := cli.VerifyCode(verifyCode)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if !result {
		t.Errorf("Verification failed")
	}
}
