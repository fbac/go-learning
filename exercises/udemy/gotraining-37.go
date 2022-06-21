// if statement
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: arggg [string]")
	}

	if os.Args[1] == "ARGGG" {
		fmt.Println("ARGGGGG!!!!")
	} else if os.Args[1] == "arggg" {
		fmt.Println("LOUDER!")
	} else {
		fmt.Println("owww..")
	}
}
