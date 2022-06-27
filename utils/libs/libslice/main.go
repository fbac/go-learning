package main

import "fmt"

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

func main() {
	intSlice := []int{0, 1, 2, 3, 4}
	float32Slice := []float32{0.0, 1.1, 2.2, 3.3, 4.4}
	float64Slice := []float64{0.0, 1.1, 2.2, 3.3, 4.4}
	stringSlice := []string{"α", "β", "γ", "δ", "ε"}

	fmt.Println(":: initial status ::")
	fmt.Println("intSlice:\t", intSlice)
	fmt.Println("float32Slice:\t", float32Slice)
	fmt.Println("float64Slice:\t", float64Slice)
	fmt.Println("stringSlice:\t", stringSlice)

	fmt.Println("\n:: PopFirst ::")
	intSlice2 := PopFirst(intSlice)
	float32Slice2 := PopFirst(float32Slice)
	float64Slice2 := PopFirst(float64Slice)
	stringSlice2 := PopFirst(stringSlice)
	fmt.Println("intSlice:\t", intSlice2)
	fmt.Println("float32Slice:\t", float32Slice2)
	fmt.Println("float64Slice:\t", float64Slice2)
	fmt.Println("stringSlice:\t", stringSlice2)

	fmt.Println("\n:: PopLast ::")
	intSlice3 := PopLast(intSlice)
	float32Slice3 := PopLast(float32Slice)
	float64Slice3 := PopLast(float64Slice)
	stringSlice3 := PopLast(stringSlice)
	fmt.Println("intSlice:\t", intSlice3)
	fmt.Println("float32Slice:\t", float32Slice3)
	fmt.Println("float64Slice:\t", float64Slice3)
	fmt.Println("stringSlice:\t", stringSlice3)

	fmt.Println("\n:: PopN(1) ::")
	intSlice4 := PopN(intSlice, 1)
	float32Slice4 := PopN(float32Slice, 1)
	float64Slice4 := PopN(float64Slice, 1)
	stringSlice4 := PopN(stringSlice, 1)
	fmt.Println("intSlice:\t", intSlice4)
	fmt.Println("float32Slice:\t", float32Slice4)
	fmt.Println("float64Slice:\t", float64Slice4)
	fmt.Println("stringSlice:\t", stringSlice4)

	fmt.Println("\n:: GetFirstN(2) ::")
	intSlice5 := GetFirstN(intSlice, 2)
	float32Slice5 := GetFirstN(float32Slice, 2)
	float64Slice5 := GetFirstN(float64Slice, 2)
	stringSlice5 := GetFirstN(stringSlice, 2)
	fmt.Println("intSlice:\t", intSlice5)
	fmt.Println("float32Slice:\t", float32Slice5)
	fmt.Println("float64Slice:\t", float64Slice5)
	fmt.Println("stringSlice:\t", stringSlice5)

	fmt.Println("\n:: GetLastN(2) ::")
	intSlice6 := GetLastN(intSlice, 2)
	float32Slice6 := GetLastN(float32Slice, 2)
	float64Slice6 := GetLastN(float64Slice, 2)
	stringSlice6 := GetLastN(stringSlice, 2)
	fmt.Println("intSlice:\t", intSlice6)
	fmt.Println("float32Slice:\t", float32Slice6)
	fmt.Println("float64Slice:\t", float64Slice6)
	fmt.Println("stringSlice:\t", stringSlice6)

}
