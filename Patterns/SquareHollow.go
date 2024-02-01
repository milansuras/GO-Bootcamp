package main

import "fmt"

func main(){

	fmt.Println("========== SQUARE HOLLOW PATTERN ========")

	for i:=1; i<=7; i++{
		for j:=1; j<=7; j++{
			if ( i==1 || j==1 || i==7 || j==7 ){
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

