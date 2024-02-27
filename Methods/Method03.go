package main

import "fmt"

type number int

func (n number) isPrime() bool {
	for i := 2; i < int(n); i++ {
		if int(n)%i == 0 {
			return false
		}
	}
	return true
}

func (n number) isDivisible(i number) bool {
	return (n % i) == 0
}

func (n *number) incrementBy(i number) {
	*n = *n + i
}
func main() {
	var n number = 7
	fmt.Println(n.isPrime())
	fmt.Println(n.isDivisible(3))
	n.incrementBy(5)
	fmt.Println(n)

}
