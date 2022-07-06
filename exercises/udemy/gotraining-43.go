/*
	Using code from gotraining-42.go:
	- slice the slice into three new slices :)
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
}
