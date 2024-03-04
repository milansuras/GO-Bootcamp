package main

import "fmt"

func main() {
	fmt.Println("Initialization of Go map")
	var m1 = map[string]int{
		"a": 7,
		"b": 10,
		"c": 1,
		"d": 5,
		"e": 9,
	}

	fmt.Println(m1)
}
