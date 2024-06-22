package service

import (
	"fmt"
	"testing"
)

func TestSvcImpl_ReqQueryPayment(t *testing.T) {
	cli, _ := GetAdminClient()
	GetResp, _ := cli.ReqQueryPayment()
	fmt.Println(GetResp.Msg)
}

func TestSvcImpl_ReqPay(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqPay(&Payment{
		Id:      "7a9b4f0a-8105-4fa0-b4ca-781d477eea0e",
		OrderId: "683479c1-757a-4a2c-8dfb-28ab474ee7ad",
		Price:   "19",
		UserId:  "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}

func TestSvcImpl_ReqAddMoney(t *testing.T) {
	cli, _ := GetAdminClient()
	UpdateResp, err := cli.ReqPay(&Payment{
		Id:      "7a9b4f0a-8105-4fa0-b4ca-781d477eea0e",
		OrderId: "c157ad12-9f53-466c-9d54-924d547ad224",
		Price:   "21",
		UserId:  "1084611818gg",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(UpdateResp.Msg)
}
