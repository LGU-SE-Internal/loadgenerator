package service

import (
	"testing"
)

func TestVerifyCodeService_VerifyCode(t *testing.T) {
	cli, _ := GetAdminClient()
	verifyCode := generateVerifyCode()
	result, err := cli.VerifyCode(verifyCode)
	if err != nil {
		t.Errorf("Request failed, err %s", err)
	}
	if !result {
		t.Errorf("Verification failed")
	}
	t.Logf("Verification code verified. The result is %v and verifyCode: %v", result, verifyCode)
}
