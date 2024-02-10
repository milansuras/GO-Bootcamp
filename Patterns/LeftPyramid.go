package main

import "fmt"

func main(){
	fmt.Println("================= LEFT PYRAMID PATTERN ====================")

	for i:=1; i<=7; i++{
		for k:=1; k<=7-i; k++{
			fmt.Print(" ")
		}
		for j:=1; j<=i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
