package main

import "fmt"

func main() {
	//declaring a variable
	var a int
	fmt.Println(a) // by default it will print zero

	//re declaring
	a = 1
	fmt.Println(a)

	//redeclaring concept in short variable declaration

	b := 0
	fmt.Println(b)

	//b :=2 it will give an error. in order to redeclare a short variable at least a new variable should introduced
	fmt.Println(b)

	b, x := 3, 4
	fmt.Println(b, x)
}
