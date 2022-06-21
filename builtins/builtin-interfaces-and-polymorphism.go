package main

import "fmt"

const aboutMsg = `
	# interfaces and polymorphism
	https://go.dev/ref/spec#Interface_types

	examples using go stdlib in examples-interfaces/

	# Polymorphism
	[!!!] In go, a value can be of MORE THAN ONE type
	If we implement certain interface with our type
	our type will *also be* of the interface type
	`

// define custom type
type potion struct {
	name        string
	description string
	action      string
	potency     int
	quantity    int
}

// define another custom type
type beverage struct {
	name        string
	description string
	action      string
	quantity    int
}

// define interface
// basically this is behavior
// for all kind of beverages
// that implement any of these methods
type alchemy interface {
	Create() (string, int)
	List() int
	Use() string
}

// potion implements alchemy interface
func (p potion) Use() string {
	return fmt.Sprintf(p.action)
}

func (p potion) List() int {
	return p.quantity
}

func (p potion) Create() (string, int) {
	return "created succesfully!", p.quantity + 1
}

// beverage implements alchemy interface
func (b beverage) Use() string {
	return fmt.Sprintf("this beverage makes you feel %s", b.action)
}

func (b beverage) List() int {
	return b.quantity
}

func (b beverage) Create() (string, int) {
	return "alcoholic beverage created succesfully!", b.quantity + 1
}

// assertions funcs
// special switch based on types implementing the interface
func GetBeverageType(a alchemy) {
	switch a.(type) {
	case potion:
		fmt.Printf("this is a potion... %s \n", a.(potion).name)
	case beverage:
		fmt.Printf("this is an alcoholic drink... %s\n", a.(beverage).name)
	}
}

func main() {

	hP := potion{
		name:        "health potion",
		description: "basic health potion",
		quantity:    1,
	}
	result, quantity := hP.Create()
	fmt.Printf("%s %s\nyou have now %v\n", hP.name, result, quantity)

	wine := beverage{
		name:        "old red wine",
		description: "an old wine forgotten in a cellar",
		action:      "dizzy",
		quantity:    1,
	}
	fmt.Println(wine.Use())

	// call assertion func
	GetBeverageType(hP)
	GetBeverageType(wine)
}
