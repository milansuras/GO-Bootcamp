package main

import "fmt"

func main(){

	var sum int
        var r int
	fmt.Println("Enter the range to find the odd number between them:-")
        fmt.Scan(&r)

	for i:=1; i<=r; i+=2{
		sum+=i
	}
	fmt.Println("The sum of odd numbers between ", "1", "to", r , "=",sum)
}

