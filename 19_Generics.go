package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Generics in go is same as Method Overloading
// type Num interface {
// 	int | int8 | int16 | float32 | float64
// }

// You can use any data type now with just one function name, makes it very readable

// To make it even better instead of creating a Num interface we can use constraints library

func add[T constraints.Ordered](a, b T) T {
	return a + b
}

func Generics() {
	fmt.Println("Generics in Go!")
	fmt.Printf("%v", add(2.1, 3.1))
}
