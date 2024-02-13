package main

import "fmt"

func main() {
	var slice = make([]int, 7, 10)
	fmt.Println(slice)
	fmt.Println("capacity of slice ", cap(slice))
	fmt.Println("length of the slice", len(slice))
}
