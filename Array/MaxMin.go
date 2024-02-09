package main

import (
	"fmt"
)

func Maxi(arr [7]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

	}
	return max
}

func Min(arr [7]int) int {
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	var arr [7]int
	fmt.Println("Enter elements in an array:=")

	for i := 0; i < len(arr); i++ {
		fmt.Println("Enter element at the index [", i, "] =")
		fmt.Scan(&arr[i])
	}

	fmt.Println("Array:=", arr)

	var maxi int = Maxi(arr)
	fmt.Println("Maximum element in the array:=", maxi)

	var min int = Min(arr)
	fmt.Println("Minimum element in the array:=", min)
}
