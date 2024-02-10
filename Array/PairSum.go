package main

import "fmt"

func pairSum(arr []int, size, target int) int {
	ans := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if arr[i]+arr[j] == target {
				ans++
			}
		}

	}
	return ans
}

func main() {

	fmt.Println("Enter the size to build array:=")
	var size int
	fmt.Scan(&size)

	var target int
	fmt.Println("Enter the target to get matching pair sum in array:=")
	fmt.Scan(&target)

	arr := make([]int, size)

	fmt.Println("Enter elements in the array:=")
	for i := 0; i < size; i++ {
		fmt.Print("Enter element at the index[", i, "]=")
		fmt.Scan(&arr[i])
	}

	fmt.Println("Array:=", arr)

	result := pairSum(arr, size, target)
	fmt.Println("Number of  pair sum in the array which is equal to the given target ==", result)
}
