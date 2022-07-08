package main

import (
	"fmt"
)

// Solution has to return N in reverse order
// without any leading 0's
//
// Example:
// 	54321 -> 12345
// 	10 -> 1
func ReverseInt(N int) {
	var enable_print int
	enable_print = N % 10

	for N > 0 {
		if enable_print == 0 && N%10 == 0 {
			enable_print = 1
		} else if enable_print >= 1 {
			fmt.Printf("%d", N%10)
		}
		N = N / 10
	}
}

func main() {
	ReverseInt(54321)
	fmt.Println()
	ReverseInt(10011)
	fmt.Println()
	ReverseInt(1)
	fmt.Println()
	ReverseInt(14)
	fmt.Println()
	ReverseInt(25)
	fmt.Println()
	ReverseInt(10)
	fmt.Println()
	ReverseInt(001)
	fmt.Println()
	ReverseInt(100)
	fmt.Println()
	ReverseInt(200)
	fmt.Println()
	ReverseInt(69000)
	fmt.Println()
}
