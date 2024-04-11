package main

import "fmt"





func main(){

	a:= []int{1,2,3,4,5,6,7,7}
	removeDuplicates(a)
	fmt.Println(a)
}


func removeDuplicates(a [] int)int{

             i:=0

	     for j:=1; j<len(a); j++{
		     if a[i] != a[j]{
			     i++
			     a[i] = a[j]
		     }
	     }
              return i+1
}
