package main

import "fmt"

func main() {
	fmt.Println("============= NUMER INCREASING REVERSE PYRAMID ==============")

	for i := 1; i <= 7; i++ {
		for j := 1; j <= 7+1-i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
