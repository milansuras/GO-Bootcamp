package main

import "fmt"

var a int = 1

//b:=2 // we cannot declare a short variable outside main function

func main() {
	fmt.Println(a)
	//fmt.Println(b)
	var x int = 11

	{
		x := 12
		fmt.Println(x)
		a = 3
		fmt.Println(a)

		var c int = 4
		fmt.Println(c)

		d := 10
		fmt.Println(d)
	}
	fmt.Println(a)
	//fmt.Println(d)
	//fmt.Println(c) // both c  and d will be  avaialable in that block

	fmt.Println(x)
}
