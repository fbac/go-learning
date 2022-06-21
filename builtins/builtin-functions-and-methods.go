package main

import (
	"fmt"
)

const aboutMsg = `
	# functions
	
	# syntax 
	func (T receiver-type) myFunc(input1 type, input2 type, inputN type) (return/s) { code }

	# syntax with variadic parameters
	# only final parameter can be variadic
	func (T receiver-type) myFunc(input1 type, input2 type, input1variadic ...type) (return/s) { code }
	
	# docs
	https://go.dev/ref/spec#Function_types

	# misc
	- in go everything is passed by value
	`

// create our struct
type potion struct {
	name   string
	effect string
}

// create methods for our struct potions
func (p potion) GetName() string {
	return fmt.Sprintf(p.name)
}

func (p potion) GetEffect() string {
	return fmt.Sprintf(p.effect)
}

// create a func that takes a bunch of potions
// function with variadic parameters
func ListPotionEffects(p ...potion) map[string]string {
	var s = make(map[string]string)
	for _, value := range p {
		s[value.name] = value.effect
	}
	return s
}

func main() {
	fmt.Println(aboutMsg)

	// initialize our potion
	// get its name via method GetName
	healthP := potion{name: "health potion", effect: "health"}
	fmt.Println("normal func:", healthP.GetName())

	// add more potions
	poisonP := potion{name: "poison potion", effect: "poisons target"}
	agilityP := potion{name: "agility potion", effect: "boost agility"}
	strengthP := potion{name: "strength potion", effect: "boost strength"}

	// defer function()
	// calling functions at the closure of current function
	// in this case, when main finalizes
	defer fmt.Println("deferred func:", poisonP.GetEffect())

	// calling variadic functions
	// ideally use a slice like s:= ListPotionEffects([]potions...)
	s := ListPotionEffects(poisonP, agilityP, strengthP)
	for key, value := range s {
		fmt.Println("variadic func\n", "\tname:", key, "\n\teffect:", value)
	}
}
