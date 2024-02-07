package main

import "fmt"

func main(){

	var unsigned uint8
	unsigned=217
	fmt.Println(unsigned)
//	unsigned=256
	fmt.Println(unsigned)// here it will overflow 

	//typecast 

	var myint int =23000000
	fmt.Println(myint)

	myint=int(unsigned)
	fmt.Println(myint)
	fmt.Println(unsigned)

	var Byte byte = 'A'
	fmt.Println(Byte)
	// here it will print ASCII of that

	Byte=255
	fmt.Println(Byte)

}
