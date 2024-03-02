package main

import "fmt"

func main(){
	// A defer statement postpones the execution of a function and pushes a function
	//call onto a list until the surronding function returns , either normally or through
	//a panic
	
	defer fmt.Println("Go")
	fmt.Println("Language")
}
