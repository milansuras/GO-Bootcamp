package main

import "fmt"

func removeEven(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			count++
		}

	}

	brr := make([]int, count)
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			brr[index] = arr[i]
			index++
		}
	}
	return brr
}

func main() {
	fmt.Println("Removing even integers from array")
	var arr []int = []int{3, 5, 7, 9, 2, 4, 6, 8, 10}
	fmt.Println("Array :=", arr)

	result := removeEven(arr)
	fmt.Println("After removing even integer new array:=", result)

}
