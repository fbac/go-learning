package main

import "fmt"

const (
	aboutMsg = `
	# Structures
	https://go.dev/ref/spec#Struct_types
	`
)

type potion struct {
	name     string
	metadata potionMetadata
}

type potionMetadata struct {
	level       string
	hp          int
	uses        int
	description string
	effect      string
}

type potionLevel map[int]string

var potionLevels potionLevel = potionLevel{
	0: "trash",
	1: "common",
	2: "rare",
	3: "epic",
	4: "legendary",
}

func main() {
	// about
	fmt.Println(aboutMsg)

	lesserHP := potion{
		name: "Lesser Health",
		metadata: potionMetadata{
			level:       potionLevels[0],
			hp:          5,
			uses:        1,
			description: "a common health potion",
			effect:      "you feel slightly restored",
		},
	}

	fmt.Printf("Lesser Health: %+v\n", lesserHP)

	//gleimsito := potion{"Gleimsito", potionMetadata{}}
	//fmt.Println(gleimsito)
}
