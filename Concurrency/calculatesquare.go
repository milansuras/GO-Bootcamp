package main

import (
    "fmt"
    "time"
)

func calculateSquare(i int) {
    time.Sleep(1 * time.Second)
    fmt.Println(i * i)
}

func main() {
    start := time.Now()
    for i := 1; i <= 100000; i++ {
        calculateSquare(i)
    }

    elapsed := time.Since(start)
    fmt.Printf("Total time taken: %s\n", elapsed)
}

