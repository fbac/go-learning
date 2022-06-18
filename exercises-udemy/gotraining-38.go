// switch with no expression specified
package main

import (
	"fmt"
	"os"
)

func main() {
	switch {
	case len(os.Args) <= 1:
		fmt.Println("No option specified...")
		fallthrough
	case len(os.Args) <= 1:
		fmt.Println("just why?")
	case len(os.Args) >= 1:
		fmt.Println("Sure, let me read your args")
		fallthrough
	case true:
		fmt.Println("Loading...")
		fallthrough
	case true:
		fmt.Println("just wait")
		fallthrough
	case true:
		fmt.Println("a little bit more")
		fallthrough
	case true:
		fmt.Println("here are your args: ", os.Args[1:])
	}
}
