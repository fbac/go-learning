/*
	Optimize the struct unstruct:
	- Save at least 4 bytes of memory
*/
package main

import (
	"fmt"
	"unsafe"
)

type unstruct struct {
	id1       string
	somebool  bool
	random    int
	somebool2 bool
	id2       string
	somemap   map[int]bool
}

type unstructStructured struct {
	id1       string
	id2       string
	somemap   map[int]bool
	random    int
	somebool  bool
	somebool2 bool
}

func main() {
	unstruct := unstruct{}
	fmt.Printf("Type:\t%T\n", unstruct)
	fmt.Printf("unstruct total size:\t\t%v\n", unsafe.Sizeof(unstruct))
	fmt.Printf("unstruct.id1 size:\t\t%v\n", unsafe.Sizeof(unstruct.id1))
	fmt.Printf("unstruct.somebool size:\t\t%v\n", unsafe.Sizeof(unstruct.somebool))
	fmt.Printf("unstruct.random size:\t\t%v\n", unsafe.Sizeof(unstruct.random))
	fmt.Printf("unstruct.somebool2 size:\t%v\n", unsafe.Sizeof(unstruct.somebool2))
	fmt.Printf("unstruct.id2 size:\t\t%v\n", unsafe.Sizeof(unstruct.id2))
	fmt.Printf("unstruct.somemap size:\t\t%v\n", unsafe.Sizeof(unstruct.somemap))

	unstructStructured := unstructStructured{}
	fmt.Printf("\nType:\t%T\n", unstructStructured)
	fmt.Printf("unstruct total size:\t\t%v\n", unsafe.Sizeof(unstructStructured))
	fmt.Printf("unstruct.id1 size:\t\t%v\n", unsafe.Sizeof(unstructStructured.id1))
	fmt.Printf("unstruct.id2 size:\t\t%v\n", unsafe.Sizeof(unstructStructured.id2))
	fmt.Printf("unstruct.somemap size:\t\t%v\n", unsafe.Sizeof(unstructStructured.somemap))
	fmt.Printf("unstruct.random size:\t\t%v\n", unsafe.Sizeof(unstructStructured.random))
	fmt.Printf("unstruct.somebool size:\t\t%v\n", unsafe.Sizeof(unstructStructured.somebool))
	fmt.Printf("unstruct.somebool2 size:\t%v\n", unsafe.Sizeof(unstructStructured.somebool2))
	fmt.Printf("\nunstruct optimized by 8 bytes!\n")
}
