package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
func GenerateSeatNumber() int {
	// 初始化随机数生成器
	return rand.Intn(30)
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

// compareOrders 比较两个 Order 实例，除了 Id 和 DifferenceMoney 字段外的所有字段是否相等。
func compareOrders(o1, o2 *Order) bool {
	equal := true

	if o1.AccountId != o2.AccountId {
		fmt.Printf("AccountId differs: %s != %s\n", o1.AccountId, o2.AccountId)
		equal = false
	}
	if o1.BoughtDate != o2.BoughtDate {
		fmt.Printf("BoughtDate differs: %s != %s\n", o1.BoughtDate, o2.BoughtDate)
		equal = false
	}
	if o1.CoachNumber != o2.CoachNumber {
		fmt.Printf("CoachNumber differs: %d != %d\n", o1.CoachNumber, o2.CoachNumber)
		equal = false
	}
	if o1.ContactsDocumentNumber != o2.ContactsDocumentNumber {
		fmt.Printf("ContactsDocumentNumber differs: %s != %s\n", o1.ContactsDocumentNumber, o2.ContactsDocumentNumber)
		equal = false
	}
	if o1.ContactsName != o2.ContactsName {
		fmt.Printf("ContactsName differs: %s != %s\n", o1.ContactsName, o2.ContactsName)
		equal = false
	}
	if o1.DocumentType != o2.DocumentType {
		fmt.Printf("DocumentType differs: %d != %d\n", o1.DocumentType, o2.DocumentType)
		equal = false
	}
	if o1.From != o2.From {
		fmt.Printf("From differs: %s != %s\n", o1.From, o2.From)
		equal = false
	}
	if o1.Price != o2.Price {
		fmt.Printf("Price differs: %s != %s\n", o1.Price, o2.Price)
		equal = false
	}
	if o1.SeatClass != o2.SeatClass {
		fmt.Printf("SeatClass differs: %d != %d\n", o1.SeatClass, o2.SeatClass)
		equal = false
	}
	if o1.SeatNumber != o2.SeatNumber {
		fmt.Printf("SeatNumber differs: %d != %d\n", o1.SeatNumber, o2.SeatNumber)
		equal = false
	}
	if o1.Status != o2.Status {
		fmt.Printf("Status differs: %d != %d\n", o1.Status, o2.Status)
		equal = false
	}
	if o1.To != o2.To {
		fmt.Printf("To differs: %s != %s\n", o1.To, o2.To)
		equal = false
	}
	if o1.TrainNumber != o2.TrainNumber {
		fmt.Printf("TrainNumber differs: %s != %s\n", o1.TrainNumber, o2.TrainNumber)
		equal = false
	}
	if o1.TravelDate != o2.TravelDate {
		fmt.Printf("TravelDate differs: %s != %s\n", o1.TravelDate, o2.TravelDate)
		equal = false
	}
	if o1.TravelTime != o2.TravelTime {
		fmt.Printf("TravelTime differs: %s != %s\n", o1.TravelTime, o2.TravelTime)
		equal = false
	}

	return equal
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
	SeatNumber             int    `json:"seatNumber"`
	Status                 int    `json:"status"`
	To                     string `json:"to"`
	TrainNumber            string `json:"trainNumber"`
	TravelDate             string `json:"travelDate"`
	TravelTime             string `json:"travelTime"`
}

type TicketArrResp struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   []Ticket `json:"data"`
}

type TicketResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Ticket `json:"data"`
}

type OrderResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Order  `json:"data"`
}

type OrderSecurity struct {
	OrderNumInLastOneHour int `json:"orderNumInLastOneHour"`
	OrderNumOfValidOrder  int `json:"orderNumOfValidOrder"`
}

type OrderSecurityResp struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Data   OrderSecurity `json:"data"`
}

type DataStringResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   Order  `json:"data"`
}

type GetOrderPriceResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type OrderArrResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	//Data   []Order `json:"data"`
	Data []Order `json:"data"`
}

type OrderInfo struct {
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

// generateVerifyCode generates a 6-digit verification code consisting of letters and numbers.
func generateVerifyCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func generateTrainTypeName(input string) string {
	startLetter := strings.ToUpper(string(input[0]))

	var MockedTrainType string

	switch startLetter {
	case "G":
		if rand.Intn(2) == 0 {
			MockedTrainType = "GaoTieOne"
		} else {
			MockedTrainType = "GaoTieTwo"
		}
	case "Z":
		MockedTrainType = "ZhiDa"
	case "T":
		MockedTrainType = "TeKuai"
	case "K":
		MockedTrainType = "KuaiSu"
	case "D":
		MockedTrainType = "DongCheOne"
	default:
		MockedTrainType = "Unknown"
	}

	return MockedTrainType
}

// generateDocumentNumber generates a DocumentNumber with 50% probability for "DocumentNumber_One"
// and 50% probability for "DocumentNumber_Two".
func generateDocumentNumber() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	if rand.Intn(2) == 0 {
		return "DocumentNumber_One"
	} else {
		return "DocumentNumber_Two"
	}
}

func GenerateTripId() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的开头字母
	letters := []rune{'Z', 'T', 'K', 'G', 'D'}

	// 随机选择一个字母
	startLetter := letters[rand.Intn(len(letters))]

	// 生成四个随机数字
	randomNumber := rand.Intn(10000)

	// 格式化成三位数字，不足三位前面补零
	MockedTripID := fmt.Sprintf("%c%03d", startLetter, randomNumber)

	return MockedTripID
}

// toLowerCaseAndRemoveSpaces converts a given string to all lower case
// and removes all spaces.
func toLowerCaseAndRemoveSpaces(input string) string {
	lowercased := strings.ToLower(input)
	noSpaces := strings.ReplaceAll(lowercased, " ", "")
	return noSpaces
}

func GenerateTrainTypeName() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的火车类型名称
	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "DongCheOne", "ZhiDa", "TeKuai", "KuaiSu"}

	// 随机选择一个火车类型名称
	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]

	return MockedTrainTypeName
}

func getMiddleElements(input string) string {
	elements := strings.Split(input, ",")

	// If the input contains less than 3 elements, return an empty string
	if len(elements) < 3 {
		return ""
	}

	middleElements := elements[1 : len(elements)-1]
	return strings.Join(middleElements, ",")
}

func generateDescription() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number with one decimal place between 0.1 and 10.0
	randomNumber := rand.Float64()*9.9 + 0.1
	numberStr := strconv.FormatFloat(randomNumber, 'f', 1, 64)

	// Determine if 'Max' should be replaced by 'Min' with a probability of 0.3
	replaceMax := rand.Float64() < 0.3
	description := "Max"
	if replaceMax {
		description = "Min"
	}

	return fmt.Sprintf("%s in %s hour", description, numberStr)
}

func generateRandomNumberString() string {
	rand.Seed(time.Now().UnixNano())
	numberLength := 10 // Length of the number string

	// Generate a random number string of the specified length
	numberStr := ""
	for i := 0; i < numberLength; i++ {
		digit := rand.Intn(10) // Generate a random digit (0-9)
		numberStr += strconv.Itoa(digit)
	}

	return numberStr
}

func ListToString(stations []string) string {

	// Use a builder for efficient string concatenation
	var builder strings.Builder

	for i, station := range stations {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("Stations[%d] %s", i, station))
	}

	result := builder.String()
	return result
}

func IntListToString(numbers []int) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder

	for i, number := range numbers {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("Numbers[%d] %d", i, number))
	}

	result := builder.String()
	return result
}

func StringToList(input string) []string {
	// Split the input string by commas and trim any leading/trailing spaces from each element
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// generateRandomTime generates a random time in the format "HH:MM:SS".
func generateRandomTime() string {
	hour := rand.Intn(24)   // 0-23
	minute := rand.Intn(60) // 0-59
	second := rand.Intn(60) // 0-59
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

type Config struct {
	StartTime string
}

type Option func(config *Config)

func WithStartTime(startTime string) Option {
	return func(config *Config) {
		config.StartTime = startTime
	}
}
func getRandomTime(opts ...Option) string {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}

	now := time.Now()

	if config.StartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", config.StartTime)
		if err != nil {
			fmt.Println("Invalid StartTime format, using current time instead.")
		} else {
			now = startTime
			// 生成1小时到1天之后的时间
			randomHours := rand.Intn(24) + 1
			randomDate := now.Add(time.Duration(randomHours) * time.Hour)
			return randomDate.Format("2006-01-02 15:04:05")
		}
	}

	// 保持原来的逻辑，生成从今天起到未来一个月内的随机日期
	randomDays := rand.Intn(30) + 1
	randomDate := now.AddDate(0, 0, randomDays)
	return randomDate.Format("2006-01-02 15:04:05")
}

// ConvertCommaSeparatedToBracketed converts a comma-separated string to a bracketed, space-separated string
func ConvertCommaSeparatedToBracketed(input string) string {
	// 删除字符串前后的空白
	input = strings.TrimSpace(input)
	// 按逗号分隔字符串，并去除每个元素前后的空白
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	// 将分隔后的元素用空格连接，并用方括号包围
	result := fmt.Sprintf("[%s]", strings.Join(parts, " "))
	return result
}

// IntSliceToString converts a slice of integers to a bracketed, space-separated string
func IntSliceToString(ints []int) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder
	builder.WriteString("[")
	for i, val := range ints {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(fmt.Sprintf("%d", val))
	}
	builder.WriteString("]")
	return builder.String()
}

// StringSliceToString converts a slice of strings to a bracketed, space-separated string
func StringSliceToString(strs []string) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder
	builder.WriteString("[")
	for i, val := range strs {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(val)
	}
	builder.WriteString("]")
	return builder.String()
}

// RandomSelectString selects a random string from a given slice of strings
func RandomSelectString(options []string) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomIndex := rand.Intn(len(options))
	return options[randomIndex]
}
