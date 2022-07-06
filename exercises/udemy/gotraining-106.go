/*
	Write an app:
	- run a goroutine
	- put 100 numbers into a channel
	- pull them
*/
package main

import "fmt"

func main() {
	bufferCh := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			bufferCh <- i
		}
		close(bufferCh)
	}()

	for j := 0; j < 100; j++ {
		fmt.Printf("Received from bufferCh:\t%v\n", <-bufferCh)
	}
}
