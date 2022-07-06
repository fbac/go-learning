# errors in Go

## Articles

[pkg error](https://pkg.go.dev/errors)
[Error handling in Golang](https://go.dev/blog/error-handling-and-go)

## Basic concepts

> :: Go Errors Bible ::
> Go does **not** use try/catch/finally
> Try/catch/finally can mislead, and usually ends in programmers looking through much code, neglecting error checking, etc.
> **Always check for errors** immediately after calling a func that returns err!
> "Errors in Go are not speciail, just values like any other" Rob Pike

## Print and Log

```go
// These functions add \n at the end
fmt.Println()

// log includes a preformatted timestamp
log.Println()

// Fatal* calls os.Exit(1)
log.Fatalln() 

/*
    With Panic deferred functions run in reverse order
    Until it's caught by a Recover
    Or exits the main goroutine
*/
log.Panicln() 
panic()
```

## Recovering

[Article: Defer Panic() and Recover()](https://blog.golang.org/defer-panic-and-recover)

> Using a deferred recover can catch panics
> Remember that the goroutines will be closed in reverse order
> from the exact point where the panic was raised.

```go
package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
```

## Custom Errors

Implementing the interface grants us the ability to have custom errors, pass them through functions and populate with custom messages.

```go
type error interface {
    Error()) string
}
```
