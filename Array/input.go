package main

import "fmt"

func main() {
    var size int
    fmt.Println("Enter the size to build array:=")
    fmt.Scan(&size)

    arr := make([]int, size) // Creating a slice with the specified size

    fmt.Println("Enter the elements in the array:=")
    for i := 0; i < size; i++ {
        fmt.Println("Enter element at index [", i, "] =")
        fmt.Scan(&arr[i])
    }

    printArray(arr, size)
}

func printArray(arr []int, size int) {
    fmt.Println("Array:", arr[:size]) // Print only the elements up to the specified size
}

