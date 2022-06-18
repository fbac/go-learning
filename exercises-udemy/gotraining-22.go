package main

import "fmt"

func main() {
	x, y, z := 0, 1, 2

	fmt.Println("x:", x, "\ty:", y, "\tz:", z)
	fmt.Printf("Is x > y:\t%t\n", x > y)
	fmt.Printf("Is y < z:\t%t\n", y < z)
	fmt.Printf("Is y != z:\t%t\n", y != z)
	fmt.Printf("Is x <= z:\t%t\n", x <= z)
	fmt.Printf("Is y >= z:\t%t\n", y >= z)
	fmt.Printf("Is x == y:\t%t\n", x == z)

}
