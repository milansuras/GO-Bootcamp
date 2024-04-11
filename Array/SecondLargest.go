package main

import "fmt"



func main(){
       var n int
       fmt.Print("Enter the size to make an array n:-")
       fmt.Scan(&n)

        //var large int
        // var secondLarge int

       arr:=make([]int , n)

       fmt.Println("Enter the element to build array")

       for i:=0; i<n; i++{
	       fmt.Print("Enter element at the index[" , i ,"]=")
	       fmt.Scan(&arr[i])
       }
       
       //printing array 
       fmt.Println(arr)


}
