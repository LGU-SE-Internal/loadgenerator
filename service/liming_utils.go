package service

import (
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
	FoodName string `json:"foodName"`
	Price    int    `json:"price"`
}

type FoodDeliveryOrder struct {
	CreatedTime        string `json:"createdTime"`
	DeliveryFee        int    `json:"deliveryFee"`
	DeliveryTime       string `json:"deliveryTime"`
	FoodList           []Food `json:"foodList"`
	Id                 string `json:"id"`
	SeatNo             int    `json:"seatNo"`
	StationFoodStoreId string `json:"stationFoodStoreId"`
	TripId             string `json:"tripId"`
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
