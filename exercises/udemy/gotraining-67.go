/*
	Assign a func to a variable, then call that func
*/
package main

import "fmt"

var myFunc func(string) string

func main() {

	decorateFunc := func(s string) string {
		return fmt.Sprintf("[decorateFunc] string passed is: %s", s)
	}

	myFunc = decorateFunc
	fmt.Println(myFunc("calling myFunc from main"))

}
