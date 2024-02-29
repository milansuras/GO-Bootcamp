package main

import "fmt"

func replace(x int) {
	x = 20 ///here the value of x is limited to this scope
}
func main() {
	fmt.Println("Function scope of a variable")

	var x int = 10
	replace(x) //here a copy of an actual parameter is passed inside calling function
	fmt.Println(x)
}
