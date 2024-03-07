package main

import "fmt"


func main(){

	var arr [7]int
	fmt.Println(arr)


	var brr *[7]int
	fmt.Println(brr)

        var brr1 =[7]int{1,2,3,4,5,6,7}
	brr=&brr1

	fmt.Println(brr[0])

}
