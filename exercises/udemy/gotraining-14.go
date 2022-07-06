package exudemy

import "fmt"

type iAmAnInt int

var x iAmAnInt

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
