/*
	Adapt code to use commaok:

	package main

	import (
		"fmt"
	)

	func main() {
		c := make(chan int)

		v, ok :=
		fmt.Println(v, ok)

		close(c)

		v, ok =
		fmt.Println(v, ok)
	}
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	// simplest solution -not performant-
	// is to use a buffer of 1 int
	// best solution: anon func
	c <- 42
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
