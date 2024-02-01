package main

import "fmt"

func main(){

	fmt.Println("============== ZERO-ONE TRIANGLE PATTERN ===========")

	for i:=1; i<=7; i++{
		for j:=1; j<=i; j++{
			if ( (i+j) % 2 == 0){
				fmt.Print("1")
			}else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}

