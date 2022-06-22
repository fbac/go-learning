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
	
	# method
	functions attached to an specific named type
	myType.Method(parameters)

	# receivers and method set
	# non-pointer receivers can work with non-pointer and pointer values
    # pointer receivers only work with pointer values
	receiver | values
	(t T)	 | T and *T
	(t *T)	 | *T

	# docs
	https://go.dev/ref/spec#Function_types
	https://go.dev/ref/spec#Method_sets

	# misc
	- in go everything is passed by value
	`

// create custom struct
type potion struct {
	name   string
	effect string
}

// concept: methods
// create methods for struct potions
// they give behavior to types
func (p potion) GetName() string {
	return fmt.Sprintf(p.name)
}

func (p potion) GetEffect() string {
	return fmt.Sprintf(p.effect)
}

// concept: variadic parameters
// create a func that takes a bunch of potions
// variadic parameters are accepted only
// for the last parameter
func ListPotionEffects(p ...potion) map[string]string {
	var s = make(map[string]string)
	for _, value := range p {
		s[value.name] = value.effect
	}
	return s
}

// concept: callbacks
// create a function that accepts a function
// as parameter (they're 1st class objects)
func Craft(f func() string, n int) string {
	name := f()
	return fmt.Sprintf("callback function: you created %v %v", n, name)
}

func mockCraft() string {
	return fmt.Sprintf("health potion")
}

// concept: recursion
// typical factorial example
func factorial(n int) int {
	// break condition
	// as we don't want to multiply anything
	// to 0
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
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

	// concept: defer function()
	// calling functions at the closure of current function
	// in this case, when main finalizes
	defer fmt.Println("deferred func:", poisonP.GetEffect())

	// calling variadic functions
	// ideally use a slice like s:= ListPotionEffects([]potions...)
	s := ListPotionEffects(poisonP, agilityP, strengthP)
	for key, value := range s {
		fmt.Println("variadic func\n", "\tname:", key, "\n\teffect:", value)
	}

	// callbacks in action
	crafted := Craft(mockCraft, 1)
	fmt.Println(crafted)
}
