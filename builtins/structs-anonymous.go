package main

import "fmt"

const (
	aboutMsg = `
	# Anonymous structures
	https://go.dev/ref/spec#Struct_types
	`
)

func main() {
	// about
	fmt.Println(aboutMsg)

	// an anonymous struct is simply a struct without name
	// the use case could be to maintain the code clean
	// for temporal structs that we don't need as a type
	// across the program, but only in one scope
	anonPotion := struct {
		name     string
		metadata string
	}{
		name:     "anonymity potion",
		metadata: "suddenly you're anonymous",
	}

	fmt.Printf("anonPotion:\n")
	fmt.Printf("\tname:\t\t%#v\n", anonPotion.name)
	fmt.Printf("\tmetadata:\t%#v\n", anonPotion.metadata)
}
