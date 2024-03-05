package main

import "fmt"

type Student struct {
	name    string
	roll_no int
}

func main() {
	var student1 = Student{name: "MILAN", roll_no: 7}
	fmt.Println(student1)

	var my_pointer = &student1 // pointer to the struct
	fmt.Println(my_pointer)

	fmt.Println(my_pointer.name)       // accessing a struct field using a pointer
	fmt.Println((*my_pointer).roll_no) // accessing a truct field using dereferncing concept
}
