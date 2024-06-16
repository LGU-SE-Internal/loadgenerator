package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 中国省会城市的英文列表
var provincialCapitalsEN = []string{
	"Beijing", "Shanghai", "Tianjin", "Chongqing",
	"Shijiazhuang", "Taiyuan", "Hohhot", "Shenyang", "Changchun", "Harbin",
	"Nanjing", "Hangzhou", "Hefei", "Fuzhou", "Nanchang", "Jinan", "Zhengzhou", "Wuhan", "Changsha", "Guangzhou",
	"Nanning", "Haikou", "Chengdu", "Guiyang", "Kunming", "Lhasa", "Xi'an", "Lanzhou", "Xining", "Yinchuan",
	"Urumqi", "Taipei",
}

// RandomProvincialCapitalEN 随机返回一个中国省会城市的英文名称
func RandomProvincialCapitalEN() string {
	rand.Seed(time.Now().UnixNano())
	return provincialCapitalsEN[rand.Intn(len(provincialCapitalsEN))]
}

// RandomIntBetween 生成并返回两个整数之间的随机整数，包括边界值。
func RandomIntBetween(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// RandomDecimalStringBetween 生成并返回两个整数之间的一位小数形式的随机数字符串，包括边界值。
func RandomDecimalStringBetween(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min              // 生成[min, max]范围内的随机整数
	decimalValue := float64(randomInt) * 0.1             // 将整数转换为一位小数
	return strconv.FormatFloat(decimalValue, 'f', 1, 64) // 转换为一位小数的字符串形式
}

// GenerateTrainNumber 随机生成火车号次字符串。
// 火车号次的格式为一个字符（G、U、D之一）后跟三位数字。
func GenerateTrainNumber() string {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 可选的首字母集合
	firstChars := []rune{'G', 'U', 'D'}
	// 随机选择一个首字母
	firstChar := firstChars[rand.Intn(len(firstChars))]

	// 生成后续的三位数字
	var numStr string
	for i := 0; i < 3; i++ {
		numStr += fmt.Sprintf("%d", rand.Intn(10))
	}

	// 拼接首字母和数字部分
	trainNumber := string(firstChar) + numStr

	return trainNumber
}

// GenerateSeatNumber 随机生成火车座位号。
// 座位号的格式为一个字符（A、B、C、D、E之一）后跟两位数字。
func GenerateSeatNumber() string {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 可选的首字母集合
	seatChars := []rune{'A', 'B', 'C', 'D', 'E'}
	// 随机选择一个首字母
	firstChar := seatChars[rand.Intn(len(seatChars))]

	// 生成后续的两位数字
	var numStr string
	for i := 0; i < 2; i++ {
		numStr += fmt.Sprintf("%d", rand.Intn(10))
	}

	// 拼接首字母和数字部分
	seatNumber := string(firstChar) + numStr

	return seatNumber
}

// GetTrainTicketClass 随机返回高铁票等级。
// 有5%的概率返回"FirstClass"（头等座），
// 15%的概率返回"BusinessClass"（一等座），
// 剩余80%的概率返回"EconomyClass"（二等座）。
func GetTrainTicketClass() int {
	rand.Seed(time.Now().UnixNano()) // 确保每次运行时随机数种子不同

	probability := rand.Intn(100) // 生成0到99之间的随机数

	switch {
	case probability < 5:
		return 0
	case probability < 20:
		return 1
	default:
		return 2
	}
}

type Order struct {
	AccountId              string `json:"accountId"`
	BoughtDate             string `json:"boughtDate"`
	CoachNumber            int    `json:"coachNumber"`
	ContactsDocumentNumber string `json:"contactsDocumentNumber"`
	ContactsName           string `json:"contactsName"`
	DifferenceMoney        string `json:"differenceMoney"`
	DocumentType           int    `json:"documentType"`
	From                   string `json:"from"`
	Id                     string `json:"id"`
	Price                  string `json:"price"`
	SeatClass              int    `json:"seatClass"`
	SeatNumber             string `json:"seatNumber"`
	Status                 int    `json:"status"`
	To                     string `json:"to"`
	TrainNumber            string `json:"trainNumber"`
	TravelDate             string `json:"travelDate"`
	TravelTime             string `json:"travelTime"`
}

type OrderResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Order  `json:"data"`
}

type DataStringResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type GetOrderPriceResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type OrderArrResp struct {
	Status int     `json:"status"`
	Msg    string  `json:"msg"`
	Data   []Order `json:"data"`
}

type Qi struct {
	BoughtDateEnd         string `json:"boughtDateEnd"`
	BoughtDateStart       string `json:"boughtDateStart"`
	EnableBoughtDateQuery bool   `json:"enableBoughtDateQuery"`
	EnableStateQuery      bool   `json:"enableStateQuery"`
	EnableTravelDateQuery bool   `json:"enableTravelQuery"`
	LoginId               string `json:"loginId"`
	State                 int    `json:"state"`
	TravelDateEnd         string `json:"travelDateEnd"`
	TravelDateStart       string `json:"travelDateStart"`
}

type Seat struct {
	DestStation  string   `json:"destStation"`
	SeatType     int      `json:"seatType"`
	StartStation string   `json:"startStation"`
	Stations     []string `json:"stations"`
	TotalNum     int      `json:"totalNum"`
	TrainNumber  string   `json:"trainNumber"`
	TravelDate   string   `json:"travelDate"`
}
