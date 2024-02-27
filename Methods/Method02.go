package main

import "fmt"

type rectangle struct {
	length float32
	width  float32
}

func (r *rectangle) increaseLength(a float32) {
	r.length += a
}
func main() {
	rect1 := rectangle{12.3, 2.2}
	//rect1.increaseLength(1.1)
	//fmt.Println(rect1) // here we can see value not changed
	p := &rect1
	p.increaseLength(8.9)
	fmt.Println(rect1)

}
