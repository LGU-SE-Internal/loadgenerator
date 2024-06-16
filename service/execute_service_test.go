package service

import (
	"fmt"
	"testing"
)

func TestSvcImpl_ReqExecuteTicket(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqExecuteTicket("7c83f029-73ab-40e5-bb6c-a45dffaab06b")
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqCollectTicket(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqCollectTicket("7f30d9c9-bc07-4494-865e-cde5d8511b1d")
	fmt.Println(GetResp.Msg)
}
