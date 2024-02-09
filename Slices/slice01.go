package main

import "fmt"

func main(){
	// declaring slice

	var s [] int
	fmt.Println(s)

	a:= [] int {1,2,3,4,5}
	fmt.Println(a)

	// make function

	b:= make([] int,7)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println("capacity", cap(b)) // capacity of slice

	c:= make([] string, 7, 14)
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))
        
        //Append function
       
	d:=make([]int ,7)
	d=append(d,10)
	fmt.Println(d)
	fmt.Println(cap(d))

}


