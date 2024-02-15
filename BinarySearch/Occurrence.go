package main

import "fmt"

func firstOccurrence(arr []int, key int) int {
	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2
	ans := -1 // Initialize ans to -1 to handle the case when key is not found
	for start <= end {
		if key == arr[mid] {
			ans = mid
			end = mid - 1
		} else if key > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = start + (end-start)/2
	}
	return ans
}

func lastOccurrence(arr []int, key int) int {
	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2
	ans := -1 // Initialize ans to -1 to handle the case when key is not found
	for start <= end {
		if key == arr[mid] {
			ans = mid
			start = mid + 1
		} else if key > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = start + (end-start)/2
	}
	return ans
}

func main() {
	var arr []int = []int{0, 0, 0, 1, 2, 2, 2, 2}
	fmt.Println("From this array printing first and last occurrence of array")
	fmt.Println("Array:=", arr)
	key := 2
	first := firstOccurrence(arr, key)
	last := lastOccurrence(arr, key)
	fmt.Println("First Occurrence of the element at the index:", first)
	fmt.Println("Last Occurrence of the element at the index:", last)
}
