/*
	Write an app:
	- run 10 goroutines
	- each goroutine add 10 numbers to a channel
	- pull the numbers from the channel and print them
*/
package main

import "fmt"

func main() {
	bufferCh := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for x := 0; x < 10; x++ {
				bufferCh <- x
			}
		}()
	}

	for j := 0; j < 100; j++ {
		fmt.Printf("Received from bufferCh:\t%v\n", <-bufferCh)
	}
}
