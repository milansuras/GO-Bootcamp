package main

import "fmt"

func main(){

	var num int
	fmt.Println("Enter an number:-")
	fmt.Scan(&num)
        
	m:=num
	var count uint8

	for num !=0 {
		num/=10
		count++
	}

	fmt.Println("The count of entered number ",m, "=", count)
}

                      
