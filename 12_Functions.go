package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

true

func isThere(name, alphabet string) (bool, int) {
	for index, letter := range name {
		if string(letter) == alphabet {
			return true, int(index)
		}
	}
	return false, -1
}

func main() {
	fmt.Printf("Functions in Go\n")

	num1 := 2
	num2 := 3

	sum := addition(num1, num2)
	fmt.Printf("Sum is: %d\n", sum)

	// Multiple return values
	// Ex: Write a function that takes a string and returns if a value exists or not, if exists return the index too
	name := "Atharva"
	exists, index := isThere(name, "h")
	fmt.Printf("%v %v", exists, index)
}
