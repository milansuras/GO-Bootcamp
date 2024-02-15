package main

import "fmt"

// declare struct struct
type Student struct {
	firstname string
	lastname  string
	roll_no   int
}

func main() {

	var student1 Student
	student1 = Student{"Milan", "suras", 25}

	fmt.Println(student1)

	// we can access firstname of student1 using . dot operator

	fmt.Println(student1.firstname)
	fmt.Println(student1.lastname)
	fmt.Println(student1.roll_no)
}
