package exudemy

import "fmt"

func main() {
	n := 1
	fmt.Printf("Decimal:%d\nBinary:%b\nHexadecimal:%X\n", n, n, n)
	nShifted := n << 1
	fmt.Printf("Decimal:%d\nBinary:%b\nHexadecimal:%X\n", nShifted, nShifted, nShifted)
}
