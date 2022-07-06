/*
	● Create a value and assign it to a variable.
	● Print the address of that value.
*/
package exudemy

import "fmt"

func main() {
	s := "I'm a string"
	fmt.Printf("%s\n", s)
	fmt.Printf("this is my memory address %v\n", &s)
}
