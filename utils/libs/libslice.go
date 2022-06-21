/*
	libslice
	author: fbac <me@fbac.dev>

	example of slice library to add
	some missing funcionality in stdlib
*/
package libslice

const (
	beginning int = 0
)

func popFirst(slice []int) []int {
	slice = append(slice[1:])
	return slice
}

func popLast(slice []int) []int {
	slice = append(slice[0 : len(slice)-1])
	return slice
}

func popN(slice []int, n int) []int {
	slice = append(slice[0:n], slice[n+1:]...)
	return slice
}

func getFirstN(slice []int, n int) []int {
	slice = append(slice[0:n])
	return slice
}

func getLastN(slice []int, n int) []int {
	slice = append(slice[n:])
	return slice
}
