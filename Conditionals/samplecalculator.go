package main

import "fmt"

func main(){
	fmt.Print("Enter a two number to perform arthmetic operation :- ")
	var a float32
	fmt.Scan(&a)
	var b float32
	fmt.Scan(&b)
        
	fmt.Println("Choose an operator := +,  - , / , * ")
	var ch string 
	fmt.Scan(&ch)


	if(ch == "+"){
		var sum float32=a+b
		fmt.Println("The sum of ", a , "+" , b, "=" , sum)
	}else if (ch == "-"){
		var sub float32=a-b
		fmt.Println("The sub of" , a , "-", b , "=", sub)
	}else if (ch == "/"){
		var div float32=a/b
		fmt.Println("The div of", a, "/" , b , "=" , div)
	}else if (ch =="*"){
		var multi float32=a*b
		fmt.Println("The multi of " , a , "*", b , "=", multi)

	}else{
		fmt.Println("Enter a valid number and choose an appropriate operator that has given")
	}
}

