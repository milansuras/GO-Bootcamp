package main

import "fmt"

func main() {
	slice1 := []string{"Java", "Golang", "Rust", "C++", "python", "Javascript", "php"}
	// creating slice from another slice

	fmt.Println(slice1)
	fmt.Println("Dynamic type language")
	slice2 := slice1[4:]
	fmt.Println(slice2)
}
