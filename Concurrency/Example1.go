package main

import (
	"fmt"
	"time"
)

func main() {
	go f1("F1")
	go f1("F2")
	fmt.Println("Sleeping for 5 seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("Main completed")

}

func f1(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":=", i)
		time.Sleep(1 * time.Second)
	}
}
