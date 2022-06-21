package main

import "fmt"

const (
	aboutMsg = `
	# arrays
	not idiomatic in Go!
	golang's way prefer to use slices over arrays.
	
	[!] Slices are built on top of arrays.

	`
)

// declare array global
// scope: global
var global [5]string

func main() {
	// about
	fmt.Println(aboutMsg)

	// declare slices with short assignment
	fmt.Printf("array global\n")
	global[0] = "first"
	global[4] = "last"
	fmt.Printf("array global: %v\n", global)

	// length
	fmt.Printf("len(global): %v\n", len(global))

}
