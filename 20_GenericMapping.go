package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// input: [1, 2, 3, 4] (n) => n * 2
// output: [2, 4, 6, 8]

func MapValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T

	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func main() {
	fmt.Println("Generic Mappings in go!")

	result := MapValues([]float64{1.1, 2.2, 3.2}, func(n float64) float64 {
		return n * 2
	})

	fmt.Printf("%v", result)
}
