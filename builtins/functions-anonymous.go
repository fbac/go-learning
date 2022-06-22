package main

import "fmt"

const (
	aboutMsg string = `
	# anonymous func
	func(input){code}(parameters)
	var := func(input){code}

	# important
	- funcs are first class objects in go, hence:
	- they can be passed to other functions
	- they can be assigned to variables
	- they can be returned from other functions
	- by passing to a var, we can create "runner" functions:
	  functions that run functions
	`
)

// Pass anonymous function as an argument
// functions can receive functions as arguments
// as functions itself are a first class object
func myAnonRunner(someAnonFunc func(p string) string) {
	fmt.Println(someAnonFunc("[myAnonRunner]"))
}

// Simple func runner
func myAnonRunner2(someAnonFunc func() string) {
	s := someAnonFunc()
	fmt.Println(s)
}

// function that returns an anonymous function
func myAnonReturner() func() string {
	fmt.Println("[myAnonReturner executed!]")
	return func() string {
		return fmt.Sprintf("[myAnonFuncReturner!]")
	}
}

func myAnonReturner2(i int) func() (int, string) {
	fmt.Println("[myAnonReturner2 executed!]")
	return func() (int, string) {
		return i, fmt.Sprintf("[myAnonFuncReturner2!]")
	}
}

func main() {
	fmt.Println(aboutMsg)

	// simple anonymous func
	// define it and run directly
	func(s string) {
		fmt.Println(s)
	}("first anonymous func")

	// assign anonymous func to a variable
	anonFunc1 := func(p string) string {
		return p + "[anonFunc1]"
	}

	// call myAnonRunner
	// passing the func
	// as first class value
	myAnonRunner(anonFunc1)

	// even simpler anon func assigned to a var
	anonFunc2 := func() string {
		return "[anonFunc2]"
	}
	myAnonRunner2(anonFunc2)

	// myAnonReturner is a func that returns a func
	// so we can assign it to a var
	anonFunc3 := myAnonReturner()
	fmt.Printf("variable anonFunc3 is type: %T\n", anonFunc3)

	// run the returned function and print the output
	anonFunc3string := anonFunc3()
	fmt.Printf("running anonFunc3 outputs: %v\n", anonFunc3string)

	// myAnonReturner accepts an int as parameter, then
	// returns a function that has been decorated with
	// the parameters passed from main
	anonFunc4 := myAnonReturner2(0)
	fmt.Printf("variable anonFunc4 is type: %T\n", anonFunc4)

	// run the function in the var and print the output
	anonFunc4int, anonFunc4string := anonFunc4()
	fmt.Printf("running anonFunc4 outputs: %v %v\n", anonFunc4string, anonFunc4int)

	/*
		I've always thought that a common
		mistake is to defer a f.Close,
		without checking the return from Close()
		hence taking the risk of having open inodes
		across the system.
		This is one of my preferred uses for defer
		in conjunction with anon funcs

		defer func() {
			closeErr := f.Close()
			if err == nil {
				err = closeErr
			}
		}()
	*/
}
