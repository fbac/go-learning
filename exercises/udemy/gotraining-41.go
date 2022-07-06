/*
	Use composite literals to:
	- create an array with 5 values type int
	- assign values to each index position

	Range over these values and print them
	Print the type of the array
*/

package exudemy

import "fmt"

var global [5]int

func main() {
	for i := 0; i < len(global); i++ {
		global[i] = i
		fmt.Printf("assign %v value in %v position\n", i, i)
	}
	fmt.Printf("array initialized to %v\n", global)
	fmt.Printf("array type is %T\n", global)
}
