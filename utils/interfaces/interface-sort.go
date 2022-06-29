/*
	sort basic methods:
	https://pkg.go.dev/sort

	primitives for sorting slices
	and user-defined collections
*/
package main

import (
	"fmt"
	"sort"
)

/*
	custom sort examples
	https://pkg.go.dev/sort#Interface
*/

// custom type
type potion struct {
	id   int
	name string
}

// custom type to implement the interface
// this will be the custom sort
// we'll order by ID field
type ByID []potion

// interface implementation
// there are 3 methods to implement
// thag magic happens between the { ... }
func (id ByID) Len() int           { return len(id) }
func (id ByID) Less(i, j int) bool { return id[i].id < id[j].id }
func (id ByID) Swap(i, j int)      { id[i], id[j] = id[j], id[i] }

func main() {
	// initialize potions and inventory
	p1 := potion{0, "health"}
	p2 := potion{1, "poison"}
	p3 := potion{2, "strength"}
	p4 := potion{3, "agility"}
	p5 := potion{4, "defense"}
	inventory := []potion{p3, p1, p3, p2, p1, p5, p4}
	fmt.Println("orig inventory:", inventory)

	// call custom sort
	sort.Sort(ByID(inventory))
	fmt.Println("sort inventory:", inventory)
}
