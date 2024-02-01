package main

import "fmt"

func main(){

	fmt.Println("Enter the number to print multiplication table:- ")
	var n int 
	fmt.Scan(&n)
	fmt.Println("Enter the range :- ")
	var r int
	fmt.Scan(&r)

	for i:=1; i<=r; i++ {
		fmt.Println( n , "*" , i , "=" , n*i )
	}
}
