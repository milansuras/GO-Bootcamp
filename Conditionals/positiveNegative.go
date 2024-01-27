package main

import "fmt"

func main(){
	fmt.Print("Enter a number :=")
	var x int
	fmt.Scan(&x)

	if ( x>0){
		fmt.Println("Entered", x , "is an positive number")

        }else if ( x<0){
		fmt.Println("Entered", x, "is an negative number")
	
	}else{
		fmt.Println("Enter a valid number")
	}

}

