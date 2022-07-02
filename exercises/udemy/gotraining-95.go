/*
	Fix the race condition on the previous exercise
	Using atomic
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const maxThreads int = 5

var wg sync.WaitGroup
var globalCounter int64

func main() {
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func() {
			n := atomic.LoadInt64(&globalCounter)
			n++
			atomic.StoreInt64(&globalCounter, n)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(globalCounter)
}
