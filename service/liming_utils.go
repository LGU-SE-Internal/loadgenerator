package service

import (
	"fmt"
	"math/rand"
	"time"
)

type OrderVO struct {
	AccountId  string `json:"accountId"`
	ContactsId string `json:"contactsId"`
	Date       string `json:"date"`
	From       string `json:"from"`
	Price      string `json:"price"`
	SeatType   int    `json:"seatType"`
	To         string `json:"to"`
	TripId     string `json:"tripId"`
}

type Food struct {
	FoodName string  `json:"foodName"`
	Price    float64 `json:"price"`
}

type FoodDeliveryOrder struct {
	CreatedTime        string  `json:"createdTime"`
	DeliveryFee        float64 `json:"deliveryFee"`
	DeliveryTime       string  `json:"deliveryTime"`
	FoodList           []Food  `json:"foodList"`
	Id                 string  `json:"id"`
	SeatNo             int     `json:"seatNo"`
	StationFoodStoreId string  `json:"stationFoodStoreId"`
	TripId             string  `json:"tripId"`
}

type FoodDeliveryOrderResponse struct {
	Status int               `json:"status"`
	Msg    string            `json:"msg"`
	Data   FoodDeliveryOrder `json:"data"`
}

type FoodDeliveryOrderArrResponse struct {
	Status int                 `json:"status"`
	Msg    string              `json:"msg"`
	Data   []FoodDeliveryOrder `json:"data"`
}

type SeatInfo struct {
	OrderId string `json:"orderId"`
	SeatNo  int    `json:"seatNo"`
}

type TripOrderInfo struct {
	OrderId string `json:"orderId"`
	TripId  string `json:"tripId"`
}

type DeliveryInfo struct {
	DeliveryTime string `json:"deliveryTime"`
	OrderId      string `json:"orderId"`
}

type Payment struct {
	Id      string `json:"id"`
	OrderId string `json:"orderId"`
	Price   string `json:"price"`
	UserId  string `json:"userId"`
}

type PaymentResponse struct {
	Status int     `json:"status"`
	Msg    string  `json:"msg"`
	Data   Payment `json:"data"`
}

type PaymentArrResponse struct {
	Status int       `json:"status"`
	Msg    string    `json:"msg"`
	Data   []Payment `json:"data"`
}

type TripPayment struct {
	TripId  string `json:"tripId"`
	OrderId string `json:"orderId"`
	Price   string `json:"price"`
	UserId  string `json:"userId"`
}

type TripPaymentResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Id      string `json:"id"`
		OrderId string `json:"orderId"`
		UserId  string `json:"userId"`
		Price   string `json:"price"`
		Type    string `json:"type"`
	} `json:"data"`
}

type TripPaymentArrResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []TripPayment `json:"data"`
}

type AccountInfo struct {
	Money  string `json:"money"`
	UserId string `json:"userId"`
}

type AccountInfoArrResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []AccountInfo `json:"data"`
}

type Money struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	Money  string `json:"money"`
	Type   string `json:"type"`
}

type MoneyResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Money  `json:"data"`
}

type TravelQueryInfo struct {
	DepartureTime string `json:"departureTime"`
	EndPlace      string `json:"endPlace"`
	StartPlace    string `json:"startPlace"`
}

type TransferTravelQueryInfo struct {
	EndStation   string `json:"endStation"`   // 目的地站
	StartStation string `json:"startStation"` // 起始站
	TrainType    string `json:"trainType"`    // 火车类型
	TravelDate   string `json:"travelDate"`   // 出行日期
	ViaStation   string `json:"viaStation"`   // 经过站
}

type TravelAdvanceResultUnit struct {
	TripId                        string   `json:"tripId"`
	TrainTypeId                   string   `json:"trainTypeId"`
	StartStation                  string   `json:"startStation"`
	EndStation                    string   `json:"endStation"`
	StopStations                  []string `json:"stopStations"`
	PriceForSecondClassSeat       string   `json:"priceForSecondClassSeat"`
	NumberOfRestTicketSecondClass int      `json:"numberOfRestTicketSecondClass"`
	PriceForFirstClassSeat        string   `json:"priceForFirstClassSeat"`
	NumberOfRestTicketFirstClass  int      `json:"numberOfRestTicketFirstClass"`
	StartTime                     string   `json:"startTime"`
	EndTime                       string   `json:"endTime"`
}

type TravelQueryArrResponse struct {
	Status int                       `json:"status"`
	Msg    string                    `json:"msg"`
	Data   []TravelAdvanceResultUnit `json:"data"`
}

type TravelQueryResponse struct {
	Status int                     `json:"status"`
	Msg    string                  `json:"msg"`
	Data   TravelAdvanceResultUnit `json:"data"`
}

type Ticket struct {
	SeatNo       int    `json:"seatNo"`
	StartStation string `json:"startStation"`
	DestStation  string `json:"destStation"`
}

type TicketArrResponse struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   []Ticket `json:"data"`
}

type TicketOrder struct {
	Date        string `json:"date"`        // 订单日期
	Email       string `json:"email"`       // 用户邮箱
	EndPlace    string `json:"endPlace"`    // 目的地
	ID          string `json:"id"`          // 订单ID
	OrderNumber string `json:"orderNumber"` // 订单号
	Price       string `json:"price"`       // 价格
	SeatClass   string `json:"seatClass"`   // 座位等级
	SeatNumber  int    `json:"seatNumber"`  // 座位号
	SendStatus  bool   `json:"sendStatus"`  // 发送状态
	StartPlace  string `json:"startPlace"`  // 出发地
	StartTime   string `json:"startTime"`   // 出发时间
	Username    string `json:"username"`    // 用户名
}

type TicketOrderArrResponse struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   []TicketOrder `json:"data"`
}

type TicketOrderResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   TicketOrder `json:"data"`
}

// init seeds the random number generator.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomTime generates a random time in the format "HH:mm:ss".
func randomTime() string {
	hour := rand.Intn(24)   // Hours range from 0 to 23
	minute := rand.Intn(60) // Minutes range from 0 to 59
	second := rand.Intn(60) // Seconds range from 0 to 59

	// Create a time.Time with the random hour, minute, and second.
	t := time.Date(0, 1, 1, hour, minute, second, 0, time.UTC)

	// Format the time to the desired layout.
	return t.Format("15:04:05")
}

func getRandomDish() string {
	dishes := []string{
		"Spaghetti Carbonara",
		"Beef Stroganoff",
		"Chicken Tikka Masala",
		"Pizza Margherita",
		"Sushi Roll",
		"Lamb Chops",
		"Grilled Salmon",
		"Caesar Salad",
		"Pad Thai",
		"Burger with Fries",
	}

	// 设置随机数种子以确保每次运行程序时都能得到不同的结果
	rand.Seed(time.Now().UnixNano())

	// 从dishes切片中随机选择一个元素
	randomIndex := rand.Intn(len(dishes))
	return dishes[randomIndex]
}

func getRandomContact() AdminContacts {
	cli := NewSvcClients()
	_, err := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	if err != nil {
		fmt.Println(err)
	}
	contacts, _ := cli.GetAllContacts()
	// 设置随机数种子以确保每次运行程序时都能得到不同的结果
	rand.Seed(time.Now().UnixNano())

	// 从contacts.Data切片中随机选择一个元素
	randomIndex := rand.Intn(len(contacts.Data))
	return contacts.Data[randomIndex]
}

func getRandomOrder() Order {
	cli := NewSvcClients()
	_, err := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	if err != nil {
		fmt.Println(err)
	}
	orders, _ := cli.ReqFindAllOrder()
	// 设置随机数种子以确保每次运行程序时都能得到不同的结果
	rand.Seed(time.Now().UnixNano())

	// 从contacts.Data切片中随机选择一个元素
	randomIndex := rand.Intn(len(orders.Data))
	return orders.Data[randomIndex]
}

func getRandomOrder_Other() Order {
	cli := NewSvcClients()
	_, err := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "222222",
		UserName:         "admin",
		VerificationCode: "123",
	})
	if err != nil {
		fmt.Println(err)
	}
	orders, _ := cli.ReqFindAllOrderOther()
	// 设置随机数种子以确保每次运行程序时都能得到不同的结果
	rand.Seed(time.Now().UnixNano())

	// 从contacts.Data切片中随机选择一个元素
	randomIndex := rand.Intn(len(orders.Data))
	return orders.Data[randomIndex]
}

func getAdjacentDates(dateStr string) (string, string, error) {
	dateLayout := "2006-01-02"
	t, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		return "", "", err
	}

	prevDay := t.AddDate(0, 0, -1).Format(dateLayout)
	nextDay := t.AddDate(0, 0, 1).Format(dateLayout)

	return prevDay, nextDay, nil
}
