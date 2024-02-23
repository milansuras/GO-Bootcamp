package main

import "fmt"

func main() {
	var intpointer *int
	var stringpointer *string
	var boolptr *bool

	fmt.Println(intpointer)    //nil
	fmt.Println(stringpointer) //nil
	fmt.Println(boolptr)       //nill

	var age int = 10
	fmt.Println(age)
	fmt.Println(&age)

	pointer := &age
	fmt.Println("value of address age:=", *pointer)
	
}
