package main

import (
	"fmt"
	"strconv"
)

func main() {

	var intValue int8 = 125
	fmt.Println("Value before conversion: ", intValue)

	fmt.Println("converting into to float32")

	var floatValue float64 = float64(intValue)
	fmt.Println("The converted value from", intValue, "to", "float is ", floatValue)

	fmt.Println("Converting int to string")

	var stringValue string = strconv.Itoa(int(intValue))

	fmt.Println(stringValue)

}
