package main

import "fmt"

func main(){
	fmt.Println("Printing even number between ranges")

	var x int
	fmt.Println("Enter the range to print even number")
	fmt.Scan(&x)

	for i:=0; i<=x; i+=2{
		fmt.Print(i," ")
	}
}

