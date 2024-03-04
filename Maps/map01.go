package main

import "fmt"

func main() {
	// A map is an unordered collection of key value pair

	//declaring map
	var m map[string]int // map keyword map [key type] value type
	fmt.Println(m)
	if m == nil {
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map is not empty")
	}

	//adding key to a nil map but it will throw runtime error
	//m["ten"] = 10
	fmt.Println(m)

}
