package main

import "fmt"

func tripletsum(arr []int, size, target int) int {
	ans := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			for k := j + 1; k < size; k++ {
				if arr[i]+arr[j]+arr[k] == target {
					ans++
				}
			}
		}

	}
	return ans
}

func main() {
	fmt.Print("Enter the size to build array:=")
	var size int
	fmt.Scan(&size)

	var arr []int
	fmt.Print("Enter the target to get the count of triplet sum:=")
	var target int
	fmt.Scan(&target)

	arr = make([]int, size)

	fmt.Println("Enter the elements in the array:=")
	for i := 0; i < size; i++ {
		fmt.Print("Enter element at the index[", i, "] =")
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)

	result := tripletsum(arr, size, target)
	fmt.Println("The count triplet to given target:=", result)
}
