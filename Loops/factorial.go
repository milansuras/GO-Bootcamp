package main

import "fmt"

func main(){

	var num int 
	fmt.Println("Enter an number to get factorial:-")
	fmt.Scan(&num)

	var fact int=1

	for i:=1; i<=num; i++{
		fact*=i
	}

	fmt.Println("The Factorial of entered number ", num , "=", fact)
}
