package main

import "fmt"

func main(){

	fmt.Println("=============== NUMBER INCREASING PYRAMID =================")

	for i:=1; i<=7; i++{
		for j:=1; j<=i; j++{
			fmt.Print(j)
		}
		fmt.Println()
	}
}
