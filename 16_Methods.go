package main

import "fmt"

type rect struct {
	witdth, height int
}

func (r *rect) area() int {
	return r.witdth * r.height
}

func (r rect) perim() int {
	return 2*r.witdth + 2*r.height
}

func Methods() {
	fmt.Printf("Methods in go\n")
	r := rect{
		witdth: 10,
		height: 5,
	}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perim())

	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perimeter: ", rp.perim())

}
