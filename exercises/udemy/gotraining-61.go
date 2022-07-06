/*
	○ create a func with the identifier foo that returns an int
	○ create a func with the identifier bar that returns an int and a string
	○ call both funcs
	○ print out their results
*/
package exudemy

import "fmt"

func foo() int {
	return 0
}

func bar() (int, string) {
	return 1, "bar"
}

func main() {
	fooInt := foo()
	barInt, barString := bar()

	fmt.Printf("fooInt: %v, barInt: %v, barString: %s\n", fooInt, barInt, barString)
}
