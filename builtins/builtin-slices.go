package main

import (
	"fmt"
)

const (
	aboutMsg = `
	# Slices
	type: type of the slice
	length [1]: current slice length
	capacity [2]: maximum capacity of the slice

	[1] initial length will be filled with 0 if not initialized
	[2] defined optionally when slice is created with make
	in case it's defined, the slice will be augmented in
	the current amount of len(slice)

	`
)

func main() {
	// about
	fmt.Println(aboutMsg)

	// declare slices with short assignment
	fmt.Println("\nx and y created and initialized with short assignment")
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// declare with make
	// make(type, length, capacity)
	// capacity is optional, if not present it's the same as length
	// if the slice is extended more than its initial capacity
	// the underlying array will be extended the amount of its capacity
	// use make when the length is known
	fmt.Println("\nz initialized with make([]int, 2, 2)")
	z := make([]int, 1, 2)
	fmt.Printf("z: %v\n", z)
	fmt.Printf("z length: %v\n", len(z))
	fmt.Printf("z capacity: %v\n", cap(z))
	fmt.Println("\nz appended 3 values, hence surpassing max cap")
	z = append(z, 1)
	z = append(z, 2)
	z = append(z, 3)
	z = append(z, 3)
	fmt.Printf("z updated: %v\n", z)
	fmt.Printf("z length: %v\n", len(z))
	fmt.Printf("z capacity: %v\n", cap(z))
	fmt.Println("\nz substitute index 0 with z[0] = 99")
	z[0] = 99
	fmt.Printf("z updated: %v\n", z)

	// slicing a slice with operator :
	fmt.Println("\nslicing x from index 3 to the end")
	fmt.Println("slice x[3:]:", x[3:])

	// append slice to slice
	// '...' is necessary when appending a slice
	// appending an int doesn't require '...'
	xy := append(x, y...)
	fmt.Println("\nxy slice, result of x.append(x, y):", xy)

	// for old school
	fmt.Println("\nfor old school over x")
	for i := 0; i < len(x); i++ {
		fmt.Println("x.key:", i, "x.value:", x[i])
	}

	// for with range
	fmt.Println("\nfor with range over y")
	for key, value := range y {
		fmt.Println("y.key:", key, "y.value:", value)
	}

	// delete element at position 4
	// note that [:4] will get to the index 4, but 4 is NOT included
	fmt.Println("\ndelete item at index 4 in xy")
	// note that in this append we'll need the notation "..."
	// to get all the values from xy[5:]
	xy = append(xy[:4], xy[5:]...)
	fmt.Println("xy delete item at index 4:", xy)

	// multidimensional slices
	// 2 dimensions
	xy2d := [][]int{x, y}
	fmt.Println("\n2 dimension slice:", xy2d)

	// 3 dimensions
	xy3d := [][][]int{xy2d, xy2d, xy2d}
	fmt.Println("\n3 dimension slice:", xy3d)

}
