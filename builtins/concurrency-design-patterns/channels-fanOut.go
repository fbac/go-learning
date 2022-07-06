/*
	Fan Out

	Use case:
		A chunk of time consuming work to be performed in parallel.
		Do as many tasks at a time as possible.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// create channels to process work
	c1 := make(chan int)
	c2 := make(chan int)

	// populate with tasks c1
	go populate(c1)

	// the work will be processed by c2
	go fanOutIn(c1, c2)

	// print the work returned from c2
	for value := range c2 {
		fmt.Println(value)
	}
}

// populate populates c1 with tasks
// in this case it's just a series of integers to be processed
// close the channel at the end to avoid deadlocks
func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// fanOutIn synchronizes the work
func fanOutIn(c1, c2 chan int) {
	// a WaitGroup will wait for all the goroutines\
	// until all the task are finished
	var wg sync.WaitGroup

	// for every task in c1:
	// - add a goroutine to the waitfroup
	// - send the work to be done
	// - wait for it to be completed
	for taskInCh1 := range c1 {
		wg.Add(1)
		go func(task int) {
			c2 <- timeConsumingWork(task)
			wg.Done()
		}(taskInCh1)
	}
	// wait for everything to finish
	wg.Wait()
	// close c2 when the work is done
	close(c2)
}

// timeConsumingWork mocks some cpu intensive func
func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(100)
}
