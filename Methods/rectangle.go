package main

import "fmt"

type Rectangle struct{
	length int
	width int
}

func (r Rectangle) area() int{
	return r.length *r.width
}

func main(){

	rectangle1:=Rectangle{12,2}
	a:=rectangle1.area()
	fmt.Println("area of rectangle:=",a)

}
