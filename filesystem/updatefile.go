package main

import ("fmt"
"os"
)

func main(){
	file,err := os.OpenFile("example.txt",os.O_WRONLY | os.O_APPEND)
	if err !=nil{
		fmt.Println("Error opening file:",err)
		return

	}
	defer file.Close()

	if _, err:= file.WriteString("new content"); err !=nil {
		fmt.Println("Error on  writing file",err)
		return
	}

	fmt.Println("File updated")
}
