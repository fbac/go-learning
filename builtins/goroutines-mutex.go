/*
	# Mutexes

	with mutexes we control the lock/unlock
	of a resource/variable/data

	this is needed to not have data race conditions

	read docs, there are different mutex types
*/
package main

import (
	"fmt"
	"sync"
)

const maxThreads int = 5

// lock resources
var mu sync.Mutex

// wait for every goroutine to finish
var wg sync.WaitGroup
var globalCounter int

func main() {
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func() {
			// lock and unlock the increment operation
			mu.Lock()
			globalCounter++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("globalCounter without data race conditions is %v\n", globalCounter)
}
