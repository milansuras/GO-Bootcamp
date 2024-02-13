package main

import "fmt"

func binarySearch(arr []int, size, key int) int {
	start := 0
	end := size - 1

	var mid int = (start + end) / 2

	for start <= end {
		if arr[mid] == key {
			return mid
		}
		if key > mid {
			start = mid + 1
		} else {

			end = mid - 1
		}
		mid = (start + end) / 2
	}
	return -1
}

func main() {
	fmt.Println("Enter size to bulid an array:=")
	var size int
	fmt.Scan(&size)

	arr := make([]int, size)
	var key int
	fmt.Println("Enter a element to search using binary search :=")
	fmt.Scan(&key)

	fmt.Println("Enter elements in the array")

	for i := 0; i < size; i++ {
		fmt.Print("Enter element at the index[", i, "]=")
		fmt.Scan(&arr[i])
	}

	fmt.Println("Created Array :=", arr)

	result := binarySearch(arr, size, key)
	fmt.Println("Element found at the index", result)
}
