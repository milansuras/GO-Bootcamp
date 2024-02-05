package main

import "fmt"

func main(){
         fmt.Println("GO programming about multiple return")
	var a , b int
	fmt.Print("Enter a & b:= ")
	fmt.Scan(&a,&b)

	var add ,  sub = addsub(a,b)
        fmt.Println("The addition of " ,a , "+", b, "=",add)
	fmt.Println("The subtraction of" , a, "-",b,"=",sub)

}

func addsub ( a,b int) (int , int) {
	return a+b, a-b
}



