package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length float32
	width  float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Triangle struct {
	base   float32
	height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.base * t.height
}

func main() {
	rect1 := Rectangle{10, 12}
	circ1 := Circle{7.7}
	tri := Triangle{18.2, 7.9}

	fmt.Println("Area of circle :=", circ1.Area())
	fmt.Println("radius of circle is :=", circ1.radius)

	fmt.Println("Area of rectangle :=", rect1.Area())
	fmt.Println("Area of triangle :=", tri.Area())

}
