package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	//AND OPERATOR
	if a != b && a <= b { //both condition should be true
		fmt.Println(true)
	}

	// OR

	if a != b || a <= b {
		fmt.Println(true) //either one of the condition to be true
	}

	//NOT
	if !(a == b) {
		fmt.Println(true)
	}
}
