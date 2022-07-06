/*
	gen generates data until context is cancelled
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	// gen will break at n = 5
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			fmt.Println("n is 5. Breaking for.")
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
