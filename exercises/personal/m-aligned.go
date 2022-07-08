package expersonal

import (
	"fmt"
	"math"
)

// The function receives two parameters:
//
// 1. string of n integers in the range [-1000:1000] represent a line of points, where:
// index is the number of the point, A[index] is the coordinate of that point
// Example: A[0] is point 0, with value N (actual value inside that field)
//
// 2. an integer M
//
// A point is M-aligned if the distance between two points in the line is divisible by M.
// The function has to return the size of the largest M-aligned subset of N points.
//
// Given: A = [0, 2]; M = 2
// The distance between 0 -> 2 is 2, the distance 2 -> 0 is 2 (distance is measured in absolute values)
// but they are the same point, hence we return 1.
func main() {
	//A := []int{-3, -2, 1, 0, 8, 7, 1}
	A := []int{7, 1, 11, 8, 4, 10}
	//A := []int{0, 2}
	M := 8

	count := 0

	for i := 0; i <= len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			fmt.Printf("A[i]: %d A[j]: %d   distance: %v\n", A[i], A[j], math.Abs(float64(A[i])-float64(A[j])))
			if int(math.Abs(float64(A[i])-float64(A[j])))%M == 0 {
				fmt.Printf("MATCH! (%d, %d) = %v\n\tA[i]: %v, A[j]: %v\n", i, j, math.Abs(float64(A[i])-float64(A[j])), A[i], A[j])
				count++
			}
		}
	}
	fmt.Println("\nTotal: ", count)
}
