package main

import "fmt"

func even(arr []int) int {
	var count int
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			count++
		}
	}
	return count
}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)

	result := even(arr)
	fmt.Println("No of odd elements in the array:=", result)
}
