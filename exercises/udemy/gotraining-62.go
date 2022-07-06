/*
● create a func with the identifier foo that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in
*/
package exudemy

import "fmt"

func foo(i ...int) int {
	var sum int
	for _, value := range i {
		sum += value
	}
	return sum
}

func bar(i []int) int {
	var sum int
	for _, value := range i {
		sum += value
	}
	return sum
}

func main() {
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nFoo := foo(n...)
	nBar := bar(n)
	fmt.Printf("foo(i ...int):\t%v\nbar(i []int):\t%v\n", nFoo, nBar)
}
