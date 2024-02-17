package main

import (
	"fmt"
	"sort"
)

func main() {

	slice1 := []string{"Zeta", "Omega", "Theta", "Kappa", "Alpha"}
	slice2 := []int{3, 1, 7, 8, 9, 5}

	fmt.Println("Before Sorting slice1:=", slice1)
	fmt.Println("Before sorting slice2:=", slice2)

	sort.Strings(slice1)
	sort.Ints(slice2)

	fmt.Println("After sorting slice1:=", slice1)
	fmt.Println("After sorting slice2:=", slice2)
}
