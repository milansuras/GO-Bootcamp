package main

import "fmt"

type Person struct {
	name string
	age  int
	Address
}

type Address struct {
	city     string
	house_no int
}

func main() {
	person1 := Person{"MILAN SURAS", 77, Address{"Vattiyoorkavu", 777}}
	fmt.Println(person1)
	fmt.Println(person1.house_no)
}
