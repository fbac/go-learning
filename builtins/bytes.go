package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sep := ""
	a := []string{"hello", "world", "goodbye", "world"}

	n := len(sep) * (len(a) - 1)
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	fmt.Println(n)
	fmt.Printf("n size: %v\n", unsafe.Sizeof(n))

	b := make([]byte, n)
	fmt.Printf("b size: %v\n", unsafe.Sizeof(b))

}
