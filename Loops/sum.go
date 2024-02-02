package main

import "fmt"

func main(){
	var sum int 
        var r int
	fmt.Println("Enter the range upto find the sum in for loop :-")
	fmt.Scan(&r)

	for i:=1; i<=r; i++{
		sum+=i
	}
	fmt.Println("The sum which ranges from", 1 , "to" , r , "=" , sum)
}

