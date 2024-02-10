package main

import "fmt"

func linearSearch(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {

		}
	}
	return true
}

func main() {

	var arr []int
	var key int
	fmt.Println("Enter an element to serach for in an array:=")
	fmt.Scan(&key)
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Array:=", arr)

	fmt.Println("Element found :=", linearSearch(arr, key))

}
