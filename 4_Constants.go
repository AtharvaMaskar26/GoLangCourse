package main

import (
	"fmt"
	"math"
)

func constants() {
	const s string = "constant"

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
