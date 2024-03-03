package main

import "fmt"

func main() {
    arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}

