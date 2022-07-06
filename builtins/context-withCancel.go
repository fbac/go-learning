package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create a Background context
	ctx := context.Background()
	fmt.Printf("ctx := context.Background\n")
	fmt.Printf("context:\t%v\n", ctx)
	fmt.Printf("context err:\t%v\n", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)

	// mutate it into context with cancellation
	// in this case we are not making use of the cancel() func
	ctx, _ = context.WithCancel(ctx)
	fmt.Printf("\nctx, _ = context.WithCancel(ctx)\n")
	fmt.Printf("context:\t%v\n", ctx)
	fmt.Printf("context err:\t%v\n", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)

	// mutate the context into a context with cancel
	// and making use of the cancel func
	ctx, cancel := context.WithCancel(ctx)
	fmt.Printf("\nctx, cancel := context.WithCancel(ctx)\n")
	fmt.Printf("context:\t%v\n", ctx)
	fmt.Printf("context err:\t%v\n", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Printf("cancel func:\t%v\n", cancel)
	fmt.Printf("cancel func type:\t%T\n", cancel)

	// call cancel
	fmt.Printf("\nCalling cancel()\n")
	cancel()
	fmt.Printf("context:\t%v\n", ctx)
	fmt.Printf("context err:\t%v\n", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Printf("cancel func:\t%v\n", cancel)
	fmt.Printf("cancel func type:\t%T\n", cancel)

	/*
		Example
	*/

	// create context with cancel
	fmt.Println("\n[Example of ctx.WithCancel] ctx.WithCancel created")
	ctx, cancel = context.WithCancel(context.Background())

	// launch goroutine
	// it will loop forever until ctx is cancelled
	fmt.Println("loop forever until ctx is cancelled")
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("go func processing data:", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("calling cancel()")
	cancel()
	fmt.Println("ctx cancelled, numbers will stop appearing.")
}
