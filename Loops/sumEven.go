package main

import  "fmt"

func main(){

	var sum int
	var r int

	fmt.Println("Enter an number to find the sum of even number between ranges: ")
	fmt.Scan(&r)

	for i:=0; i<=r; i+=2{

		sum+=i
	}

	fmt.Println("The sum of even number between ranges =",sum)
}


