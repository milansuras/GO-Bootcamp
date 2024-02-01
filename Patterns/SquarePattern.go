package main

import "fmt"

func main(){

	fmt.Print("====================== SQUARE PATTERN =======================")
         


	for i:=1; i<=7; i++{
		for j:=1; j<=7; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
