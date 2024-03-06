package main

import "fmt"

type Executor interface{
	Execute()
}

type Thread struct{
}


func (t Thread) Execute(){
	fmt.Println("Executing thread")
}

func main(){
	var exe Executor
	exe=Thread{}
	exe.Execute()
}
