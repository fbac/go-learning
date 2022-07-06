/*
	Pull the values off the channel using a select:

	package main

	import (
		"fmt"
	)

	func main() {
		q := make(chan int)
		c := gen(q)

		receive(c, q)

		fmt.Println("about to exit")
	}

	func gen(q <-chan int) <-chan int {
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
	q := make(chan int)
	c := gen(q)
	receive(c, q)
	fmt.Println("Shutdown.")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		// send 100 integers to c
		for i := 0; i < 100; i++ {
			c <- i
		}
		// send exit signal and close q channel
		q <- 1
		close(q)
	}()
	return c
}

func receive(c <-chan int, q <-chan int) {
	for {
		select {
		// no need to declare vars, we simply check we receive something from ch
		case <-c:
			fmt.Printf("Received from c:\t%v\n", <-c)
		case <-q:
			fmt.Printf("Shutting down...\n")
			return
		}
	}
}
