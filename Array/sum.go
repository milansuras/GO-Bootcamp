package main

import "fmt"



func Sum(arr [7]int) int {
    sum := 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }
    return sum
}

func main() {
    fmt.Println("Calculating sum of array:=")

    var arr [7]int
    fmt.Println("Enter the elements in the array")

    for i := 0; i < len(arr); i++ {
        fmt.Println("Enter element at index [", i, "] =")
        fmt.Scan(&arr[i])
    }

    fmt.Println(arr)

    result := Sum(arr)
    fmt.Println("The sum of array :=", result)
}

