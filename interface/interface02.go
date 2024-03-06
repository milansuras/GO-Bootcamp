package main

import "fmt"

type Bird interface{
	Fly()
}

type Eagle struct{
}

func(e Eagle) Fly(){
	fmt.Println("Eagle flying over the cloud")
}

type Pigeon struct{
}

func (p Pigeon) Fly(){
	fmt.Println("Pigeon flying at an normal height")
}

type Penguin struct{
}

func (p Penguin) Fly(){
	fmt.Println("Penguin cannot fly")
}


func flynow(b Bird){
	b.Fly()


}
func main(){

	fmt.Println("Polymorphism interface as a contract")

         flynow(Eagle{})
	 flynow(Pigeon{})
         flynow(Penguin{})
}
