/*
	Use composite literals to:
	- create slice of type int
	- assign 10 values

	Range over these values and print them
	Print the type of the slice
*/

package main

import "fmt"

var slice []int

func main() {
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		fmt.Printf("assign %v value in %v position\n", i, i)
	}
	fmt.Printf("slice initialized to %v\n", slice)
	fmt.Printf("slice length is %v\n", len(slice))
	fmt.Printf("slice capacity is %v\n", cap(slice))
	fmt.Printf("slice type is %T\n", slice)
}
