package main

import "fmt"

func printArray2(arr []int) {
	fmt.Println(arr)
}

func main() {
	fmt.Println("Programme for passing array in a function")

	var arr []int
	arr = []int{2, 3, 4, 5, 6, 7, 5}

	printArray2(arr)
}
