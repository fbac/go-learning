package main

import "fmt"

func main() {
	var ptrX int
	x := &ptrX
	fmt.Printf("%T\n", x)
}
