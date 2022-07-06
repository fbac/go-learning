package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create context with timeout (50 ms)
	ctx, cancel := context.WithTimeout(context.Background(), 15000*time.Millisecond)
	defer cancel()

	select {
	// more than 1 second will trigger this cond
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		// prints "context deadline exceeded"
		fmt.Println(ctx.Err())
	}

}
