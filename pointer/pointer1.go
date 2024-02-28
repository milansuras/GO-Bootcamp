package main

import "fmt"

func main() {
	x := 7
	a := &x
	fmt.Println("Value of x:=", x)
	fmt.Println("value of at address :=", a)
	fmt.Println("Value of a:=", *a)

	fmt.Println("Modifying :=")
	*a = 20
	fmt.Println("value of x:=", x)
	fmt.Println("Value of a:=", *a)
}
