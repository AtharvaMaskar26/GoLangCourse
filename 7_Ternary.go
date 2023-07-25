package main

import (
	"fmt"
	"time"
)

func ternary() {
	fmt.Println("Welcome to the clock")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend!")
	default:
		fmt.Println("It is a weekday!")
	}

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Good Morning")
	case t >= 12:
		fmt.Println("Good Afternoon!")
	case t > 17:
		fmt.Println("Good Evening")
	}
}
