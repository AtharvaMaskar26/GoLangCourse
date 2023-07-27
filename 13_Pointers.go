package main

import (
	"fmt"
)

func callByValue(value int) {
	value = 0
}

func callByReference(value *int) {
	*value = 0
}

func Pointers() {
	fmt.Printf("Pointers in Go\n")

	// Initial Value
	value := 10

	// Calling call by value
	callByValue(value)
	fmt.Printf("Call By Value: %d\n", value) // Value here will stay 10 and not change

	// Calling by Reference
	callByReference(&value)
	fmt.Printf("Call by reference %d\n", value) // Value here should change to 0
}
