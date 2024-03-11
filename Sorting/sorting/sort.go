package main

import (
	"fmt"
)

func main() {
	// Define an array of integers
	numbers := []int{64, 34, 25, 12, 22, 11, 90}

	// Print the original array
	fmt.Println("Original array:", numbers)

	// Call the bubble sort function
	bubbleSort(numbers)

	// Print the sorted array
	fmt.Println("Sorted array:", numbers)
}

// Function to perform bubble sort on an array of integers
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

