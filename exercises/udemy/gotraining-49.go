/*
	Using code from last exercise:
	- Add a new element outside var declaration
	- Show the map looping over all the keys and all the values inside them
*/

package main

import "fmt"

func main() {

	potions := map[string][]string{
		"Health":  {"Lesser", "Normal", "Greater"},
		"Poison":  {"Lesser", "Normal", "Greater"},
		"Agility": {"Lesser", "Normal", "Greater"},
	}

	for typePotion, levelPotion := range potions {
		fmt.Printf("type: %v\tlevels: %v\n", typePotion, levelPotion)
	}

	fmt.Println()
	potions["Strength"] = []string{"Lesser", "Normal", "Greater"}
	for k, v := range potions {
		fmt.Printf("type %v:\n", k)
		for _, j := range v {
			fmt.Printf("\t%v\n", j)
		}
	}
}
