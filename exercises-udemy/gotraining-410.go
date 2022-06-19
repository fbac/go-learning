/*
	Using code from last exercise:
	- delete a record
	- print map using range
*/

package main

import "fmt"

const (
	_       = iota
	Lesser  = "Lesser"
	Normal  = "Normal"
	Greater = "Greater"
)

func main() {

	potions := map[string][]string{
		"Health":  {Lesser, Normal, Greater},
		"Poison":  {Lesser, Normal, Greater},
		"Agility": {Lesser, Normal, Greater},
	}

	for typePotion, levelPotion := range potions {
		fmt.Printf("type: %v\tlevels: %v\n", typePotion, levelPotion)
	}

	fmt.Println()
	potions["Strength"] = []string{Lesser, Normal, Greater}
	for k, v := range potions {
		fmt.Printf("type %v:\n", k)
		for _, j := range v {
			fmt.Printf("\t%v\n", j)
		}
	}

	fmt.Println()
	delete(potions, "Poison")
	for k, v := range potions {
		fmt.Printf("type %v:\n", k)
		for _, j := range v {
			fmt.Printf("\t%v\n", j)
		}
	}
}
