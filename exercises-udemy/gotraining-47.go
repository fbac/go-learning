/*
	Create a slice from multidimensional strings
*/

package main

import "fmt"

func main() {
	x := []string{"x", "y", "z"}
	y := []string{"a", "b", "c"}
	xy := [][]string{x, y}
	fmt.Println(xy)

	// old school
	// height is len(xy)
	// width is len(x) or len(y) indistinctively
	// [!] vars used innecesarely, only for learning
	height := len(xy)
	width := len(x)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Println("x:", i, "y:", j, "val:", xy[i][j])
		}
	}

	// using range makes it easier
	// range xy returns the 2 rows:
	// 0: [x y z]
	// 1: [a b c]
	// then we iterate over these rows
	// to get the columns
	// we can use valueRow (values inside each row)
	// to get the proper index
	for matrixRow, valueRow := range xy {
		fmt.Println("row:", matrixRow, valueRow)
		for matrixColumn, valueColumn := range valueRow {
			fmt.Printf("position: (%v.%v) x: %v y: %v val: %v\n", matrixRow, matrixColumn, matrixRow, matrixColumn, valueColumn)
		}
	}

}
