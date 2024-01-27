package main

import "fmt"

func main(){
	var a int 
	fmt.Println("Enter number a:-")
	fmt.Scan(&a)

	var b int
	fmt.Println("Enter number b:-")
	fmt.Scan(&b)

	sum := a+b
	fmt.Println("The sum " ,a , "+", b , "=",sum)
}