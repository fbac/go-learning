/*
	ctx.Deadline will close when the Deadline is met
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Set a Deadline
	d := time.Now().Add(2500 * time.Millisecond)

	// Create the context
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Overslept: Deadline is higher than this time.After")
	case <-ctx.Done():
		// ctx.Done() will be called after the Deadline has been met
		fmt.Println("Received ctx.Done()")
		// Print a context deadline exceeded error
		fmt.Println(ctx.Err())
	}

}
