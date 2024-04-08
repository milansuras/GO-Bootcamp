package main

import (
	"fmt"
	"container/list")
func main(){

	var intList list.List
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)

	for i:=intList.Front(); i!=nil; i=i.Next(){
		fmt.Println(i.Value.(int))
	}




}
