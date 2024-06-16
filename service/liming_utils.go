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
