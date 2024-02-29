package main

import "fmt"

func replace2(x *int) {
	*x = 20
}

func main() {
	var x int = 7
	fmt.Println("Before modificaion x:=", x)
	replace2(&x)
	fmt.Println("After modification x:=", x)
}
