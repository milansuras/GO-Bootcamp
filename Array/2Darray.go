package main

import "fmt"

func main() {
	fmt.Println("2D ARRAY ")

	var arr [3][4]int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

	fmt.Println("Array:=", arr)
}
