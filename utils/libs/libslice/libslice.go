/*
	libslice
	author: fbac <me@fbac.dev>

	example of slice library to add
	some missing funcionality in stdlib
*/
package libslice

type sliceType interface {
	int | float32 | float64 | string
}

func PopFirst[sT sliceType](slice []sT) []sT {
	slice = append(slice[1:])
	return slice
}

func PopLast[sT sliceType](slice []sT) []sT {
	slice = append(slice[:len(slice)-1])
	return slice
}

func PopN[sT sliceType](slice []sT, n int) []sT {
	slice = append(slice[:n], slice[n+1:]...)
	return slice
}

func GetFirstN[sT sliceType](slice []sT, n int) []sT {
	slice = append(slice[:n])
	return slice
}

func GetLastN[sT sliceType](slice []sT, n int) []sT {
	slice = append(slice[n+1:])
	return slice
}
