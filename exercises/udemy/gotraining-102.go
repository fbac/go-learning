/*
	Make this code work:

	package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
*/

package main

import (
	"fmt"
)

func main() {
	// we make chan bidirectional
	cs := make(chan int)

	// now data can be sent over it
	go func() {
		cs <- 42
	}()

	// we retrieve data
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
