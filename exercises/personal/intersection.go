/*
	Calculate the intersection between
	two slices: xs and ys
*/
package main

import (
	"fmt"
)

func main() {
	// find the intersection between
	// this two collections
	xs := []int{0, 1, 2, 3, 4}
	ys := []int{0, 1, 2, 5, 6}

	// hold the intersection in res
	var res []int

	// for this, we'll create a map
	// struct is used here because it's 0 bytes
	m := make(map[int]struct{})

	// initialize map as int: {empty}
	// this map keys are the values of xs
	for _, x := range xs {
		m[x] = struct{}{}
	}

	// loop over ys
	for _, y := range ys {
		// retrieve key:value, discard key
		// if value m[y] exists, means
		// that it's in xs and ys
		// add it to the intersection
		if _, ok := m[y]; ok {
			res = append(res, y)
		}
	}

	fmt.Println("xs :", xs)
	fmt.Println("ys :", ys)
	fmt.Println("m :", m)
	fmt.Println("res :", res)
}
