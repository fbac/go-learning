package main

import (
	"fmt"
	"sort"
)

type potion struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ByName []potion

func (name ByName) Len() int           { return len(name) }
func (name ByName) Less(i, j int) bool { return name[i].Name < name[j].Name }
func (name ByName) Swap(i, j int)      { name[i], name[j] = name[j], name[i] }

type ByID []potion

func (id ByID) Len() int           { return len(id) }
func (id ByID) Less(i, j int) bool { return id[i].Id < id[j].Id }
func (id ByID) Swap(i, j int)      { id[i], id[j] = id[j], id[i] }

func main() {
	p1 := potion{Id: 0, Name: "health"}
	p2 := potion{Id: 1, Name: "strength"}
	p3 := potion{Id: 2, Name: "agility"}
	p4 := potion{Id: 3, Name: "invisibility"}
	p5 := potion{Id: 4, Name: "teleport"}
	p6 := potion{Id: 5, Name: "poison"}
	p7 := potion{Id: 6, Name: "shrink"}

	potions := []potion{p1, p3, p5, p7, p1, p5, p6, p4, p3, p2, p5, p7}
	fmt.Println("unsorted inventory:\t", potions)

	sort.Sort(ByID(potions))
	fmt.Println("sort by ID:\t", potions)

	sort.Sort(ByName(potions))
	fmt.Println("sort by Name:\t", potions)
}
