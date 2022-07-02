/*
	# goroutines

	they need some sort of sync and waiting
	that can be done with waitgroups and channels

	otherwise the goroutines will continue running
	and closed when main finalizes
*/
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("[goroutine %v] this won't be printed\n", i)
		}()
	}
}
