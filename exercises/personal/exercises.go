package expersonal

import (
	"fmt"
	"math"
	"sort"
)

/*
A zero-indexed array A consisting of N integers is given. Rotation of the array means that each element is
shifted right by one index, and the last element of the array is also moved to the first place.
*/
func CyclicRotate(A []int, K int) []int {
	if K == 0 || len(A) == 0 || len(A) == 1 {
		return A
	}

	fmt.Println("\n\tlen(A):", len(A))
	fmt.Println("\tK before modulus:", K)
	// this modulo op controls the rotation
	// in case len(A) < K
	if len(A) < K {
		K = K % len(A)
	}
	fmt.Println("\tK after modulus:", K)

	newStart := A[len(A)-K:]
	newEnd := A[:len(A)-K]
	fmt.Println("\tnewStart", newStart)
	fmt.Println("\tnewEnd", newEnd)

	return append(newStart, A[:len(A)-K]...)
}

/* Odd Element in Array */
func isOdd(n int) bool {
	return n%2 == 1
}

func OddOccurences(A []int) (oddOne int) {
	var histogram = make(map[int]int)

	// Populate histogram with values in A
	// and the time it appears in A (hence the ++)
	for _, value := range A {
		histogram[value]++
	}

	// look for the odd one
	for k, value := range histogram {
		if isOdd(value) {
			oddOne = k
			return oddOne
		}
	}

	return 0
}

/*
	binary search - nails
*/

func nails(a []int, b []int, c []int) int {
	n := len(a)   // a length
	m := len(c)   // c length - max nails to use
	minNails := 1 // min nails will always be 1
	maxNails := m
	var mid int
	var missing bool

	total := -1

	// ?
	maxCoord := m*2 + 1
	nailed := make([]int, maxCoord)

	for minNails <= maxNails {
		missing = false
		mid = (maxNails + minNails) / 2

		for i := 0; i < maxCoord; i++ {
			nailed[i] = 0
		}

		for i := 0; i < mid; i++ {
			nailed[c[i]]++
		}

		for i := 1; i < maxCoord; i++ {
			nailed[i] = nailed[i] + nailed[i-1]
		}

		for i := 0; i < n; i++ {
			if nailed[a[i]-1] == nailed[b[i]] {
				missing = true
			}
		}

		if missing {
			minNails = mid + 1
		} else {
			maxNails = mid - 1
			total = mid
		}
	}

	return total
}

/*
	frog
*/

func FrogJmp(X int, Y int, D int) int {
	initialPos := float64(Y - X)
	nJumps := initialPos / float64(D)
	return int(math.Ceil(nJumps))
	// return int(math.Ceil(float64(float64(Y-X) / float64(D))))
}

/*
	array missing element
*/

func Missing(A []int) int {
	// reduce complexity
	// allows to use our algorithm
	sort.Ints(A)

	// best case scenario
	if len(A) == 0 {
		return 1
	}

	// iterate over keys
	// check the next
	// if the next natural number is missing
	// that's what we have to return
	for key, value := range A {
		if value+1 != A[key+1] {
			return value + 1
		}
	}

	// program shouldn't get here
	// this would mean an error
	return -1
}

/*
	Numbers on a tape
*/

func NumbersOnATape(A []int) int {
	// Best case scenario
	if len(A) == 2 {
		return int(math.Abs(float64(A[0] - A[1])))
	}

	// Total array value
	totalValue := 0
	for _, k := range A {
		totalValue = totalValue + k
	}

	relativePos := A[0]                      // we start calculating from A[0] and we'll be moving the position
	initialValue := totalValue - relativePos // initialValue is the totalValue of the rest of the array
	// hold the min integer in this var
	// we have to initialize it at max int in order to not
	// interfere in the calculations
	// as each N is in the range [-1000 ... 1000]
	minInt := math.MaxInt64
	fmt.Println(minInt)

	// now we simply iterate the array and check the current min
	// start at i = 1, as our relativePos is in A[0]
	for i := 1; i < len(A)-1; i++ {
		difference := int(math.Abs(float64(relativePos) - float64(initialValue)))

		if difference < minInt {
			minInt = difference
		}

		relativePos = relativePos + A[i] // increment the relative position with the value of the current element
		totalValue = totalValue - A[i]   // decrement the totalValue as we'll be moving up the array
	}
	return minInt
}

func main() {

	fmt.Println(":: Array :: Cyclic Rotate")
	A := []int{0, 1, 2}
	K := 2
	S := CyclicRotate(A, K)
	fmt.Println("\n\tA:", A)
	fmt.Println("\tS:", S)

	fmt.Println("\n:: Array :: Odd Occurences")
	A = []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	i := OddOccurences(A)
	fmt.Println("\n\tA:", A)
	fmt.Println("\ti:", i)

	fmt.Println("\n:: Array :: Find Missing Int")
	A = []int{2, 3, 1, 5}
	i = Missing(A)
	fmt.Println("\n\tA:", A)
	fmt.Println("\ti:", i)

	fmt.Println("\n:: Array :: NumbersOnATape")
	//A = []int{3, 1, 2, 4, 3}
	A = []int{-1, 8, 8, 2, 998, 445}
	i = NumbersOnATape(A)
	fmt.Println("\n\tA:", A)
	fmt.Println("\ti:", i)
}
