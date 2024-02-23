package main

import "fmt"

func increment(a int) {
	a = a + 2

}

func increment2(x *int) {
	*x++
}
func main() {
	var a int = 10
	fmt.Println(a)

	increment(a)
	fmt.Println(a)

	increment2(&a)
	fmt.Println(a)
}
