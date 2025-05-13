package main

import "fmt"

var a int8 = 23

func main() {

	fmt.Println(a)

	{
		a++
		fmt.Println(a)
	}

	var b int8 = 27

	{
		b++
		fmt.Println(b)
	}

	fmt.Println(b)
	fmt.Println(a)

	{
		var c int8 = 29
		fmt.Println(c)
		c++
	}

	//fmt.Println(c) // this will give an error as that c is declared in that particular block and it will be avialble in that block only.

}
