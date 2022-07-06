/*

	Make this code work

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
*/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// 1. using go func
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// 2. using a buffered channel
	c2 := make(chan int, 1)
	c2 <- 42
	fmt.Println(<-c2)
}
