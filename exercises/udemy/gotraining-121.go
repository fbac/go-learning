/*
	Create a package:
	- export a function
	- document it
	- check it via web
*/
package exudemy

import "fmt"

// Exported shows some cool stuff
func Exported() {
	fmt.Println("😎")
}

// main runs Exported once
func main() {
	Exported()
}
