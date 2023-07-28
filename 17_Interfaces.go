package main

// You might have observed that you cannot use name of a function more than once, interface solves this.
// It is similar to method overloading in OOPs

import (
	"fmt"
	"math"
)

// Interfaces are named collection of methods signatures
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	fmt.Println("Interfaces in Go!")

	r := rect{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}

	measure(&r)
	measure(&c)

}
