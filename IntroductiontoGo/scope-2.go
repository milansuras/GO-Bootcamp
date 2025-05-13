package main

import "fmt"

func main() {

	x := 10
	fmt.Println("1. Value of x in main scope:", x)

	{
		fmt.Println("2. Value of x inside first block (before shadowing):", x)

		x := 20
		fmt.Println("3. Value of x inside first block (after shadowing):", x)

		{
			fmt.Println("4. Value of x inside second block (before shadowing):", x)

			x := 30
			fmt.Println("5. Value of x inside second block (after shadowing):", x)

			x = 35
			fmt.Println("6. Modified value of x inside second block:", x)
		}

		fmt.Println("7. Value of x in first block after exiting second block:", x)

		fmt.Println("8. Modified value of x in first block:", x)
	}

	fmt.Println("9. Value of x in main scope after exiting all blocks:", x)

	if true {

		x := 50
		fmt.Println("10. Value of x inside if statement:", x)
	}

	fmt.Println("11. Value of x in main scope after if statement:", x)

	for i := 0; i < 3; i++ {

		x := 100 + i
		fmt.Println("12. Value of x inside for loop:", x)
	}

	fmt.Println("13. Value of x in main scope after for loop:", x)

	x = 15
	fmt.Println("14. Final value of x in main scope:", x)
}
