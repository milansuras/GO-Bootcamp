package main

import "fmt"

func Occurrences(arr [7]int, x int) int {
	var count int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			count++
		}
	}
	return count
}

func main() {
	var arr [7]int
	fmt.Println("Enter elements in the array:=")
	var x int

	for i := 0; i < len(arr); i++ {
		fmt.Println("Enter Element at the index[", i, "]=")
		fmt.Scan(&arr[i])

	}
	fmt.Print("Enter a number to count the occurences:=")
	fmt.Scan(&x)
	fmt.Println(arr)

	result := Occurrences(arr, x)
	fmt.Println("Number of occurences:=", result)

}
