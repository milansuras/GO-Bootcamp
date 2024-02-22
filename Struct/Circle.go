package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := Circle{radius: 10}
	fmt.Println(c1)
	fmt.Println("radius:=", c1.radius)
	fmt.Println("Area of circle:=", c1.area())
}
