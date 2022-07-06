package main

import (
	"context"
	"fmt"
)

func main() {
	// Create the context withValue
	// it will call ctx.Done() when the value is found
	ctx := context.WithValue(context.Background(), "mockValue1", "mockValue2")
	foo(ctx, "mockValue1")

	ctx = context.WithValue(ctx, "mockValue3", "mockValue4")
	foo(ctx, "mockValue3")

	foo(ctx, "mockValue0")
}

func foo(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}
