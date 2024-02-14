package main

import "fmt"

func swap(x, y int) (int, int) {
    return y, x
}

func reverseArray(arr []int, size int) {
    var start, end int = 0, size-1

    for start < end {
        arr[start], arr[end] = swap(arr[start], arr[end])
        start++
        end--
    }
}

func main() {
    var size int
    fmt.Print("Enter the size of array to build := ")
    fmt.Scan(&size)

    arr := make([]int, size)
    fmt.Println("Enter elements in the array:")

    for i := 0; i < size; i++ {
        fmt.Printf("Enter element at index [%d]: ", i)
        fmt.Scan(&arr[i])
    }

    // printing array
    fmt.Println("Original array:", arr)

    reverseArray(arr, size)

    // printing reversed array
    fmt.Println("Reversed array:", arr)
}

