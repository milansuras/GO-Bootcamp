package main

import "fmt"

func main(){

    fmt.Println("Continue")
    for i:=1; i<10; i++{
      defer fmt.Println(i)
  }
  fmt.Println("completed")
}