/*
	Context
		Read carefully:
		https://pkg.go.dev/context
		https://go.dev/blog/context
		https://go.dev/blog/pipelines
		https://peter.bourgon.org/blog/2016/07/11/context.html


	Use case
		context groups goroutines under the same context,
		so when it's closed all the goroutines are finished.
		This behavior prevents the program from leaking goroutines.
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	// Background is non-nil and empty, never canceled, no values, no deadline.
	// Typically used as main/base ctx.
	ctx := context.Background()
	fmt.Printf("ctx := context.Background")
	fmt.Printf("context:\t%v\n", ctx)
	fmt.Printf("context err:\t%v\n", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
}
