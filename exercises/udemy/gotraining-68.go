/*
	● Create a func which returns a func
	● assign the returned func to a variable
	● call the returned func
*/
package main

import "fmt"

func myFuncReturner(s string) func() {
	return func() {
		fmt.Printf(s)
	}
}

func main() {
	myFuncRunner := myFuncReturner("calling myFuncReturner from main.myFuncRunner\n")
	myFuncRunner()
}
