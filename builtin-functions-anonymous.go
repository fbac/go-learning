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
		return fmt.Sprintf("[myAnonFuncReturned!]")
	}
}

func main() {
	fmt.Println(aboutMsg)

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

	// simpler anon func
	anonFunc2 := func() string {
		return "[anonFunc2]"
	}
	myAnonRunner2(anonFunc2)

	// assign a returned func to a value
	anonFunc3 := myAnonReturner()
	fmt.Printf("variable anonFunc3 is type: %T\n", anonFunc3)

	// run the function and print the output
	anonFunc3string := anonFunc3()
	fmt.Printf("running anonFunc3 outputs: %v\n", anonFunc3string)

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
