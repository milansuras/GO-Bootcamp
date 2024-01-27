package main

import "fmt"

func main(){
	fmt.Print("Enter a number :-")
	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Entered ", num , "is an even number")
	
	}else{
		fmt.Println("Entered", num, "is an odd number")
	}
}
