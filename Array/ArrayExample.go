package main

import "fmt"

func main(){

	// array declaration

	var a[5] int 
	fmt.Println(a)

	// array declaration with elements

	var b = [5] int {1,2,3,4,5}
	fmt.Println(b)

	// short declaration
	c:= [5]int{11,12,13,14,15}
	fmt.Println(c)

	//sparse array declaration

	d:=[5] int {1:10,2,3}
	fmt.Println(d)
        

	//implicit length declaration
	e:=[...]int{3,4,5,6,7,8,9,1,1,1}
	fmt.Println(e)

	// to get the length of array using len function

	fmt.Println(len(e))
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))

	// accessing elements in array

	fmt.Println(a[1])
	fmt.Println(e[3])

}
