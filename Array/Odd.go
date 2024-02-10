package main

import "fmt"

func oddArray(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			count++
		}
	}
	return count
}
func main() {
	var arr []int

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr)

	ans := oddArray(arr)
	fmt.Println("No of odd elements in the array:=", ans)
}
