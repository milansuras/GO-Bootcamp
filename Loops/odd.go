package main

import "fmt"

func main(){

	fmt.Println("Printing odd number between the ranges")
	fmt.Print("Enter the range:=")
         var r uint8
	 fmt.Scan(&r)
         var i uint8
	 for  i=1; i<=r; i+=2{
		 fmt.Print(i , " " )
	 }
 }
