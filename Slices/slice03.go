package main

import "fmt"

func main() {
	var a []string = []string{"apple", "orange", "green_apple", "grapes"}
	fmt.Println(a)

	slice01 := a[:3]
	slice02 := a[2:]
	slice03 := a[:]

	fmt.Println(slice01)
	fmt.Println(slice02)
	fmt.Println(slice03)
}
