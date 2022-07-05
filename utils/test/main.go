package main

import "fmt"

// potion struct defines a potion
type potion struct {
	id   int
	name string
}

// NewPotion returns a newly crafted potion
func NewPotion(id int, name string) *potion {
	return &potion{id, name}
}

// String returns the potion name
func (p *potion) String() string {
	return p.name
}

// main is just a mock, in real use cases this would be
// a complete program
func main() {
	p := NewPotion(0, "health")
	fmt.Println(p.String())
}

/*
	running go test will run all the tests inside the package
 	the expected results:

	$ go test -v
	=== RUN   Test_potion_String
	=== RUN   Test_potion_String/test:_health
	=== RUN   Test_potion_String/test:_poison
	=== RUN   Test_potion_String/test:_agility
	=== RUN   Test_potion_String/test:_hd12345
	=== RUN   Test_potion_String/test:_whatever👌🏻
	--- PASS: Test_potion_String (0.00s)
		--- PASS: Test_potion_String/test:_health (0.00s)
		--- PASS: Test_potion_String/test:_poison (0.00s)
		--- PASS: Test_potion_String/test:_agility (0.00s)
		--- PASS: Test_potion_String/test:_hd12345 (0.00s)
		--- PASS: Test_potion_String/test:_whatever👌🏻 (0.00s)
	=== RUN   TestNewPotion
	--- PASS: TestNewPotion (0.00s)
	=== RUN   TestString
	--- PASS: TestString (0.00s)
	PASS
	ok      github.com/fbac/go-resources/utils/test 0.002s

*/
