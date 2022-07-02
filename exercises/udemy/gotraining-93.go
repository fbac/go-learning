/*
	Create an incrementer with goroutines
		- only one counter var
		- maxThreads const
	Each goroutine:
		- reads value into new var
		- yields processor
		- increment new var
		- write new var into counter
	Demonstrate there's a race condition
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

const maxThreads int = 5

var wg sync.WaitGroup
var globalCounter int

func main() {
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func() {
			n := globalCounter
			runtime.Gosched()
			n++
			globalCounter = n
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(globalCounter)
}

/*
	Expected output:

	$ go run -race gotraining-93.go
==================
WARNING: DATA RACE
Read at 0x0000005cb540 by goroutine 8:
  main.main.func1()
      /home/aranda/ownCloud/projects/go-resources/exercises/udemy/gotraining-93.go:29 +0x29

Previous write at 0x0000005cb540 by goroutine 7:
  main.main.func1()
      /home/aranda/ownCloud/projects/go-resources/exercises/udemy/gotraining-93.go:32 +0x46

Goroutine 8 (running) created at:
  main.main()
      /home/aranda/ownCloud/projects/go-resources/exercises/udemy/gotraining-93.go:28 +0x4a

Goroutine 7 (finished) created at:
  main.main()
      /home/aranda/ownCloud/projects/go-resources/exercises/udemy/gotraining-93.go:28 +0x4a
==================
5
Found 1 data race(s)
exit status 66
*/
