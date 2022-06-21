package main

import "fmt"

func main() {

	// switch without expression
	switch {
	// default case can be anywhere in the switch
	default:
		fmt.Println("[switch1] default")
	case false:
		fmt.Println("[switch1] this won't be printed")
	case true:
		fmt.Println("[switch1] this will be printed")
	}

	// switch with expression
	switch 1 < 2 {
	case true:
		fmt.Println("[switch2] this will be printed")
		// fallthrough statement makes the sequency
		// to check next case
		fallthrough
	case true:
		fmt.Println("[switch2] falling through: this will also be printed")
	default:
		fmt.Println("[switch2] this will never be printed")
	}
}
