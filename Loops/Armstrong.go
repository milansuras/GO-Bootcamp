package main

import "fmt"

func main(){

	var num int
	fmt.Println("Enter a number to check for Armstrong:-")
	fmt.Scan(&num)

	m:=num
        var sum int

	for num !=0 {
		rem:=num%10
		sum=sum + rem * rem * rem
		num= num/10
	}

	if sum==m {
		fmt.Println("Entered number := ", m , "Is an Armstrong number")
	}else {
		fmt.Println("Entered number :=", m , "Is not an Armstrong number")
	}
}

	
