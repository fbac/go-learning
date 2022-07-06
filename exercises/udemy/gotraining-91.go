/*
	Launch some goroutines
		Each goroutine prints out something
	use waitGroups to make sure everything finishes.
*/
package exudemy

import (
	"fmt"
	"sync"
)

const maxThreads int = 5

var wg sync.WaitGroup

func main() {
	wg.Add(maxThreads)

	for i := 0; i < maxThreads; i++ {
		go func(id int) {
			fmt.Println("worker num:", id)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

/*
	Result expected: (order may change)
	$ go run gotraining-91.go
	worker num: 2
	worker num: 0
	worker num: 4
	worker num: 3
	worker num: 1
*/
