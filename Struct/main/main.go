package main

import (
	"Struct/module"
	"fmt"
)

func main() {
	e := module.Employee{Id: 100,
		//name = "xyz,
		City: "New york",
	}
	fmt.Println("Employee Id  and city:=", e)

	//city1:=module.city{
	//city_name: "Mumbai"
	//}

	//fmmt.Println("City:=",city1} // we cannot accesss city struct  from outside package

	// Also we cannot access some fields in the Employee struct as well
	//why its because its start with small letter
	// When you dont want your struct or field accessed from any outside package write
	//in it in small letter

}
