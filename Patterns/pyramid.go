package main

import "fmt"

func main(){

	fmt.Println("=============== PYRAMID ============")

	for i:=1; i<=7; i++{
		for j:=1; j<=i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
