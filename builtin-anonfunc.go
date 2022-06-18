package main

import "fmt"

// Pass anonymous function as an argument
func myAnonRunner(someAnonFunc func(p string) string) {
	fmt.Println(someAnonFunc("namedFunc"))

}

func main() {

	// assign anonymous func to a variable
	value := func(p string) string {
		return p + "anonFunc"
	}

	// call myAnonRunner
	// passing the func 
	// as first class value
	myAnonRunner(value)
}
