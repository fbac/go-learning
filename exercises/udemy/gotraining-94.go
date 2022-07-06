/*
	Fix the race condition on the previous exercise
	Using mutex
*/
package exudemy

import (
	"fmt"
	"runtime"
	"sync"
)

const maxThreads int = 5

var wg sync.WaitGroup
var m sync.Mutex
var globalCounter int

func main() {
	wg.Add(maxThreads)
	for i := 0; i < maxThreads; i++ {
		go func() {
			m.Lock()
			n := globalCounter
			runtime.Gosched()
			n++
			globalCounter = n
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(globalCounter)
}
