package main

import "fmt"

func main() {
	//creating slice from array

	var a []int = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)

	var b []int = a[0:len(a)]
	fmt.Println(b)
}
