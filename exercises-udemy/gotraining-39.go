package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: command [your favorite sport]")
	}

	switch os.Args[1] {
	case "rugby":
		fmt.Println("much steroids bro")
	case "hockey":
		fmt.Println("nice sticks")
	case "football":
		fmt.Println("mainstream")
	case "basketball":
		fmt.Println("you like boring")
	case "tennis":
		fmt.Println("is this even a sport")
	default:
		fmt.Println("I don't even know this")
	}
}
