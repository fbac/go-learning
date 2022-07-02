/*
	Spells:
	- Create a spell struct
	- Spells has a method Use
		- it accepts a *spell
	- Magic is an interface
		- to use Magic, the spell has to habe Use method
	- func doSomething
		- spell as parameter
		- it calls Use
	- you can pass values of type *spell to doSomething
	- you cannot pass values of type spell to doSomething
*/
package main

import "fmt"

type spell struct {
	name       string
	invocation string
	effect     string
}

func (s *spell) Use() {
	fmt.Printf("%v\n%v\n", s.invocation, s.effect)
}

type magic interface {
	Use()
}

func doSomething(m magic) {
	m.Use()
}

func main() {
	s1 := spell{name: "Fireball", invocation: "Vas Flam", effect: "*you burn the target for 15 damage*"}
	doSomething(&s1)
}
