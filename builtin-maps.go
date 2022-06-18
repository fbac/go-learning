package main

import "fmt"

const (
	aboutMsg = `
	# Maps
	map[type]type

	# Example
	map[int]string
	map[1] = "One"

	`
)

// create a map with long assignment
// scope: global
var global map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

func main() {
	// about
	fmt.Println(aboutMsg)

	// print global map
	fmt.Println("global map:", global)

	// create a map with short assignment
	// scope: main
	potions := map[string]int{
		"Lesser Health":  15,
		"Health":         30,
		"Greater Health": 45,
	}

	// print map
	fmt.Println("Potions:", potions)

	// print a key
	fmt.Println("Health potion heals you for", potions["Health"], "HP")

	// keys that doesn't exist shows a value of 0
	fmt.Println("This will be a zero value as index doesn't exist:", potions["Poison"])

	// check if a key exist
	fmt.Println("\ncheck if a potion exist and its value (zero valued)")
	value, ok := potions["Poison"]
	fmt.Println("Poison potion value?", value)
	fmt.Println("Does Poison potion exist?", ok)

	// check if a key exist
	// this is called "comma ok" idiom in go
	// and a good practice to check keys
	fmt.Println("\ncomma ok check: if value, ok := potions[\"Poison\"]; ok { ... }")
	if value, ok := potions["Poison"]; ok {
		fmt.Println("if will never get in here", value)
	} else {
		fmt.Println("Poison potion doesn't exist")
	}

	// loop over a map
	fmt.Println("\nloop over a map with range")
	for key, value := range potions {
		fmt.Println(key, ":", value)
	}

	// add new elements
	fmt.Println("\nadd a new element")
	potions["Poison"] = 30
	fmt.Println("new potion: ", potions)
	fmt.Println("Poison potion does", potions["Poison"], "damage over time")

	// delete elements
	fmt.Println("\ndelete(potions, \"Poison\"), we don't like poisons\nloop over potions again")
	delete(potions, "Poison")
	for key, value := range potions {
		fmt.Println(key, ":", value)
	}
}
