package main

import ("fmt"
"sync"
"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go f1 ("F1",&wg)
	go f1 ("F2",&wg)

	fmt.Println("Main: waiting for goroutine to finish")
	wg.Wait()
	fmt.Println("Main completed")
}

func f1(name string, wg *sync.WaitGroup){
	for i:=0; i<10; i++{
		fmt.Printf("%v : i %d\n", name , i)
		time.Sleep( 1 * time.Second)
	}
	wg.Done()
}
