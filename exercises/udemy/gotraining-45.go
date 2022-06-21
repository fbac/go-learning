/*
	Create a slice:
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	Delete items in position 0, 5 and 7
	Print the slice
*/

package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("initial slice: %v\n", slice)

	// delete index 0
	slice = append(slice[1:2], slice[2:]...)
	fmt.Printf("deleted index 0 from slice: %v\n", slice)

	// delete index 5
	slice = append(slice[:5], slice[6:]...)
	fmt.Printf("deleted index 5 from slice: %v\n", slice)

	// delete index 7
	slice = append(slice[:7], slice[8:]...)
	fmt.Printf("deleted index 7 from slice: %v\n", slice)
}
