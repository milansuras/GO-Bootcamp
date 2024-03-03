package main

import "fmt"

func modifiedArray(arr []int) {
	arr[0] = 1
	fmt.Println(arr)
}

func main() {
	var arr []int = []int{2, 3, 4, 5, 6, 7}
	fmt.Println("Before modification array :=", arr)

	fmt.Println("After modification array:=")
	modifiedArray(arr)
}
