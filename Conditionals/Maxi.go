package main

import "fmt"

func main(){
	
	var a ,b,c int
	fmt.Print("Enter three number to findout maximum from them:-")
	fmt.Scan(&a,&b,&c)

	if a>=b && a>=c {
		fmt.Println(a , "Is greatest among ", b,"and",c)
	}else if b>=c{
		fmt.Println(b , "Is greatest among ", a , "and", c)
	}else if a==b && b==c && c==a {
		fmt.Println(a , b , c , "Entered number all are equal")
	}else {
		fmt.Println(c , "Is greatest among ", a, "and", b)
	}
}
