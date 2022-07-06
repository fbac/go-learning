/*
	Create a map:
	- key string
	- value array of strings
	- Store three records
	- Print them referencing them by index pos
*/

package exudemy

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
}
