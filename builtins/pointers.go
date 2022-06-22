package main

import "fmt"

const (
	aboutMsg string = `
	# pointers
	
	> pointers are variables that store memory addresses
	to certain variables.

	> for example, *int is a pointer that stores a memory address
	which references an integer

	> in the end, we would have:

	[mem.table]
	0x00001: pointer to 0x02000, which is an int
	...
	...
	0x02000: int value X

	> printing the address of x would yield 0x02000
	> printing the address of pX would yield 0x00001
	> printing the value of pX would yield 0x02000
	`
)

func main() {
	fmt.Println(aboutMsg)

	var vInt int
	var pInt *int
	fmt.Printf("[1] create vInt (integer) and pInt (pointer to integer)\n")
	fmt.Printf("\n# vInt\n")
	fmt.Printf("vInt value:\t %v\n", vInt)
	fmt.Printf("vInt address:\t %v\n", &vInt)
	fmt.Printf("vInt type:\t %T\n", vInt)
	fmt.Printf("\n# pInt\n")
	fmt.Printf("pInt value:\t %v\t\t(expected nil, as pointer is not initialized)\n", pInt)
	fmt.Printf("pInt address:\t %v\t(pointers are variables, they have also an address)\n", &pInt)
	fmt.Printf("pInt type:\t %T\n", pInt)

	fmt.Printf("\n[2] assign to pInt the memory address of vInt \n")
	fmt.Printf("    > pInt = &vInt\n")
	fmt.Printf("    > we're 'referencing' the memory address of vInt from pInt\n")
	pInt = &vInt
	fmt.Printf("\n# vInt\n")
	fmt.Printf("vInt value:\t %v\n", vInt)
	fmt.Printf("vInt address:\t %v\n", &vInt)
	fmt.Printf("vInt type:\t %T\n", vInt)
	fmt.Printf("\n# pInt\n")
	fmt.Printf("pInt value:\t %v\t(expected same memory address as vInt, since we just initialized it)\n", pInt)
	fmt.Printf("*pInt value:\t %v\t\t(operator * dereferences the address, giving us the value stored in it)\n", *pInt)
	fmt.Printf("pInt address:\t %v\n", &pInt)
	fmt.Printf("pInt type:\t %T\n", pInt)
}
