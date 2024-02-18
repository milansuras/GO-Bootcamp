package main

import "fmt"

type employee struct {
	name   string
	id     int
	salary float32
}

func main() {

	var employee1 employee // zero value struct
	fmt.Println("Employee1 :=", employee1)
}
