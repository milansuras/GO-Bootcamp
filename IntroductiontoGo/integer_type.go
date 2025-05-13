package main

import "fmt"

func main() {
	// uint8   ranges from 0 to 255 contains only unsigned integers 8 bit (positive numbers)
	// uint16  ranges from 0 to 65535 contains only unsigned integers 16 bit
	// uint32  ranges from 0 to 2^32 -1
	// int8    ranges from (-128 to 127) signed integer contains negative numbers also 8 bit
	// int16   ranges from (-32768 to 32767) 16 bit
	// int32   ranges 32 bit  (-2147483648 to 2147483647)
	// int64   ranges 64 bit  (-9223372036854775808 to 9223372036854775807)

	fmt.Println("=== OVERFLOW EXAMPLES (going beyond upper limit) ===")

	// Overflow in uint8
	var a uint8 = 255
	fmt.Println("Starting with max uint8:", a)
	a++
	fmt.Println("After increment (overflow):", a) // output returns to zero
	a += 10
	fmt.Println("After adding 10:", a)
	a = 255
	fmt.Println("Reset to max uint8:", a)
	a += 10
	fmt.Println("After adding 10 (overflow):", a) // 255+10=265, but 265%256=9

	// Overflow in int8
	var b int8 = 127
	fmt.Println("\nStarting with max int8:", b)
	b++
	fmt.Println("After increment (overflow):", b) // 127+1=-128 (wraps to minimum value)
	b = 127
	fmt.Println("Reset to max int8:", b)
	b += 10
	fmt.Println("After adding 10 (overflow):", b) // 127+10=137, but overflows to -119

	fmt.Println("\n=== UNDERFLOW EXAMPLES (going below lower limit) ===")

	// Underflow in uint8
	var c uint8 = 0
	fmt.Println("Starting with min uint8:", c)
	c--
	fmt.Println("After decrement (underflow):", c) // 0-1=255 (wraps to maximum value)

	// Underflow in int8
	var d int8 = -128
	fmt.Println("\nStarting with min int8:", d)
	d--
	fmt.Println("After decrement (underflow):", d) // -128-1=127 (wraps to maximum value)
	d = -128
	fmt.Println("Reset to min int8:", d)
	d -= 10
	fmt.Println("After subtracting 10 (underflow):", d) // -128-10=-138, but underflows to 118
}
