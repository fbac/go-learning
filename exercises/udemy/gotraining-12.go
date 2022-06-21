package main

import "fmt"

var x, y, z int

func main() {

	// a
	fmt.Println(x, y, z)

	// b - zero values
	fmt.Printf("%#v", x)
	fmt.Printf("%#v", y)
	fmt.Printf("%#v", z)
}
