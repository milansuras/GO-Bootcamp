package main

import "fmt"

func main(){

	fmt.Println("STRING TYPE")

	// In Golang string can be created using double quotes"" and backticks ``

	var str string // string variable declaration by default it will print an empty string

	str="Golang \nLearning " // here \n treated as a special case after Golang. Learning will //  be printed in new line 

	fmt.Println(str)
       
	x := `Gin\nframework`
	fmt.Println(x)
}

