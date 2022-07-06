/*
	Using code from gotraining-43.go:
	- append value 42 to slice
	- print the slice
	- append 97 98 99 in one statement
	- print the slice
	- append slice2, slice3 and slice4 to slice
	- print the resulting slice
*/

package exudemy

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

	fmt.Printf("let's slice the slice\n")
	slice2 := slice[:2]
	slice3 := slice[2:8]
	slice4 := slice[8:]
	fmt.Printf("slice2 initialized to %v\n", slice2)
	fmt.Printf("slice3 initialized to %v\n", slice3)
	fmt.Printf("slice4 initialized to %v\n", slice4)

	fmt.Printf("let's append things\n")
	slice = append(slice, 42)
	fmt.Printf("slice is now: %v\n", slice)

	slice = append(slice, 97, 98, 99)
	fmt.Printf("slice is now: %v\n", slice)

	slice = append(slice, slice2...)
	slice = append(slice, slice3...)
	slice = append(slice, slice4...)
	fmt.Printf("slice is now: %v\n", slice)
}
