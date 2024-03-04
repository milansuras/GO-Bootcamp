package main

import "fmt"

func main() {
	//creating an array using make function
	var arr = make([]int, 5)
	fmt.Println("Array:=", arr)
	fmt.Println("Length :=", len(arr))
	fmt.Println("Capacity of array:=", cap(arr))
}
