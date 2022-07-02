/*
	# atomic

	operations with atomic
	instead of mutex, perform atomic ops
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const maxThreads int = 5

var wg sync.WaitGroup
var globalCounter int32

func main() {
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func() {
			n := atomic.LoadInt32(&globalCounter)
			n++
			atomic.StoreInt32(&globalCounter, n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("globalCounter:", globalCounter)
}
