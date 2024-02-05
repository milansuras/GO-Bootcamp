package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter number a and b:-")
	fmt.Scan(&a, &b)

	var result int = sum(a, b)
	fmt.Println("The sum of ", a, "+", b, "=", result)
}

func sum(a int, b int) int {
	return a + b
}
