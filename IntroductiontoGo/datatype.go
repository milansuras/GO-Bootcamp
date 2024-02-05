package main

import "fmt"

func main(){
       var name string // by default it will be an empty string ""
	   fmt.Println(name)
	   
	   name="GOLANG"
	   fmt.Println(name)

	   var a int // By default it will print zero
	   fmt.Println(a)
	   a=777
	   fmt.Println(a)

	   var b bool // By default it will print false
	   fmt.Println(b)
 

	   // re declaration concept

	   x:= 10
//	   x:=20

	   fmt.Println(x) // here it will print nothing we cannot redecalre but using short varible
	   // in-order o declare atleast a new variable should be introduced

	   y:=777
	   y, z:= 9,10

	   fmt.Println(y,z)

}

