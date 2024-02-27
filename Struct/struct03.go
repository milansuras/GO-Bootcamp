package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	student1 := student{id: 7, name: "YXZ"}
	student2 := student1

	if student1 == student2 {
		fmt.Println("Both structs are same")

	} else {
		fmt.Println("Both are not same")
	}
}
