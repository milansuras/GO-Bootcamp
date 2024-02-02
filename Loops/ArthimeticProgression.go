package main

import "fmt"

func main(){
	var a , x , d int
	fmt.Println("Enter the First term (a), range (x) difference(d) :- ")
	fmt.Scan(&a)
	fmt.Scan(&x)
	fmt.Scan(&d)
        var next_term int=a
	for i:=1; i<=x; i++{
		fmt.Print(next_term, " ")
		next_term+=d
	}
}

