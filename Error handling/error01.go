package main

import (
	"errors"
	"fmt"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("Program for error handling")
	fmt.Print("Enter  number a:=")
	var a float32
	fmt.Scan(&a)
	fmt.Print("Enter number b:=")
	var b float32
	fmt.Scan(&b)

	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("RESULT:=", result)
}
