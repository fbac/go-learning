/*
	Create a func which encloses the scope of a var
*/
package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x value in main:\t", x)

	func() {
		x := 1
		fmt.Println("x value in f:\t\t", x)
	}()
}
