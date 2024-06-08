package service

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestSvcImpl_ReqSeatCreate(t *testing.T) {
	cli, _ := GetBasicClient()
	resp, err := cli.ReqSeatCreate(&SeatCreateInfoReq{
		TravelDate:  "2024-06-06 14:16:00",
		TrainNumber: "777",
		DestStation: "shenzhen",
		SeatType:    7,
		TotalNum:    8,
		Stations:    []string{"shenzhen", "suzhou", "hong kong"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("create response: %+v", resp)

	tranNumber := ""
	randn := rand.Int() % 15
	if randn > 7 {
		tranNumber = "G" + strconv.Itoa(rand.Int()%100)
	} else if randn > 3 {
		tranNumber = "D" + strconv.Itoa(rand.Int()%100)
	} else {
		tranNumber = strconv.Itoa(rand.Int() % 100)
	}
	t.Logf("tranNumber: %s", tranNumber)

	respGetTicket, err := cli.ReqGetTicketLeft(&SeatCreateInfoReq{
		TravelDate:  "2024-06-06 14:16:00",
		TrainNumber: tranNumber,
		DestStation: "shenzhen",
		SeatType:    7,
		TotalNum:    8,
		Stations:    []string{"shenzhen", "suzhou", "hong kong"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("respGetTicket: %+v", respGetTicket)
	if respGetTicket.Status != 1 {
		t.Error("respGetTicket.Status != 1")
	}
	if respGetTicket.Msg != "Get Left Ticket of Internal Success" {
		t.Error("respGetTicket.Data != Get Left Ticket of Internal Success", respGetTicket)
	}
}
