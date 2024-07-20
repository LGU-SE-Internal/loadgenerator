package service

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
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

func TestSvcImpl_ReqSeatCreate_v2(t *testing.T) {
	cli, _ := GetBasicClient()
	var seatSvc SeatService = cli

	randomOrder := getRandomOrder()
	var stations []string
	totalNum := rand.Intn(10)

	// 设置随机数种子，以确保每次运行程序时都能得到不同的随机数
	rand.Seed(time.Now().UnixNano())

	// 生成一个[0, 1)之间的浮点数
	randomFloat := rand.Float64()

	// 如果随机数小于0.5，则执行if代码块；否则，执行else代码块
	if randomFloat < 0.5 {
		stations = []string{randomOrder.To, randomOrder.From}
	} else {
		stations = []string{randomOrder.From, randomOrder.To}
	}

	seatCreateInfoReq := &SeatCreateInfoReq{
		TravelDate:  randomOrder.TravelDate,
		TrainNumber: randomOrder.TrainNumber[1:],
		DestStation: randomOrder.To,
		SeatType:    randomOrder.SeatClass,
		TotalNum:    totalNum,
		Stations:    stations,
	}

	resp, err := seatSvc.ReqSeatCreate(seatCreateInfoReq)

	if err != nil {
		t.Error(err)
	}
	if resp.Data.DestStation != randomOrder.To {
		t.Error("respGetTicket.Data.DestStation != randomOrder.To")
	}

	respGetTicket, err := cli.ReqGetTicketLeft(seatCreateInfoReq)
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
