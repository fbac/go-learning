/*
	Semaphore
*/
package main

import "fmt"

func main() {
	semaphore := make(chan int)
	stop := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			semaphore <- i
		}
		stop <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			semaphore <- i
		}
		stop <- true
	}()

	go func() {
		<-stop
		<-stop
		close(semaphore)
	}()

	for n := range semaphore {
		fmt.Printf("Semaphore:\t%v\n", n)
	}
}
