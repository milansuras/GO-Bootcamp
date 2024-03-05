package main

import "fmt"

type city struct {
	city1 string
	city2 string
	city3 string
}

func main() {
	// Go programme to explain struct are value type
	c1 := city{"trivandrum", "pattom", "vattapara"}
	c2 := c1 //A copy of struct c1 is passed to c2

	fmt.Println("city_list:=", c1)
	fmt.Println("city_lsit:=", c2)

	c2.city2 = "Bengaluru"
	c2.city3 = "Delhi"

	fmt.Println("city:=", c1)
	fmt.Println("city:=", c2)
}
