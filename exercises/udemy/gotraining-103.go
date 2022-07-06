/*
	Starting code:
	- Retrieve the data from gen correctly

	package main

	import (
		"fmt"
	)

	func main() {
		c := gen()
		receive(c)

		fmt.Println("about to exit")
	}

	func gen() <-chan int {
		c := make(chan int)

		for i := 0; i < 100; i++ {
			c <- i
		}

		return c
	}
*/
package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)
	fmt.Println("Shutdown.")
}

func receive(c <-chan int) {
	for value := range c {
		fmt.Printf("Received from channel:\t%v\n", value)
	}
}

func gen() chan int {
	c := make(chan int)

	// enclose the for and close(c) in a func
	// now gen and receive will run at the same time
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
