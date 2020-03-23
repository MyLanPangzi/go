package main

import (
	"fmt"
	"time"
)

func getDay(day time.Weekday) time.Weekday {
	fmt.Println(day)
	return day
}
func main() {
	fmt.Println("When's Saturday?")
	switch today := time.Now().Weekday(); today {
	case getDay(time.Saturday):
		fmt.Println("Saturday.")
	case getDay(time.Friday):
		fmt.Println("Friday")
	default:
		fmt.Println("Working day...")
	}
}
