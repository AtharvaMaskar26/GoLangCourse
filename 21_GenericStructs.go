package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	name string
	Data T
}

func main() {
	fmt.Println("Generic Structs in Go")
	u := User[int]{
		ID:   0,
		name: "Atharva",
		Data: 4,
	}

	fmt.Printf("%v", u)
}
