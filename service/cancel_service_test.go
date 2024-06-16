package service

import (
	"fmt"
	"testing"
)

func TestSvcImpl_ReqCalculate(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqCalculate("790bcfd5-82d2-4717-aa9f-e00bef992268")
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqCancelTicket(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqCancelTicket("790bcfd5-82d2-4717-aa9f-e00bef992268", "test1")
	fmt.Println(GetResp.Msg)
}
