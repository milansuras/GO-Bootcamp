package main

import "fmt"

func main(){
	fmt.Println("================ NUMBER CHANGING PYRAMID ================")

	var count int=1 

	for i:=1; i<=7; i++{
		for j:=1; j<=i; j++{
			fmt.Print(count," ")
			count++
		}
		fmt.Println()
	}
}
