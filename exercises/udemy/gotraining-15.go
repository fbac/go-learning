package main

import "fmt"

type iAmAnInt int

var x iAmAnInt
var y iAmAnInt

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 1
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
