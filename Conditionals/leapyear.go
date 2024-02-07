package main

import "fmt"

func main() {
    var year int
    fmt.Println("Enter a year to check for leap year or not:")
    fmt.Scan(&year)

    if year%4 == 0 {
        if year%100 == 0 {
            if year%400 == 0 {
                fmt.Println("Entered year:", year, "is a leap year")
            } else {
                fmt.Println("Entered year:", year, "is not a leap year")
            }
        } else {
            fmt.Println("Entered year:", year, "is a leap year")
        }
    } else {
        fmt.Println("Entered year:", year, "is not a leap year")
    }
}

