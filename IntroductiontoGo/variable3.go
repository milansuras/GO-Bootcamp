package main

import "fmt"

var x int =10 // it is a package level variable can access through out the package


func main(){
      fmt.Println(x)

      function()

      // in Golang variable can be grouped 

      var (
	      my int =7
	     name string="Golang"
      )
      
      fmt.Println(my , name)

}

func function(){
	fmt.Println("value:=" , x)
}
