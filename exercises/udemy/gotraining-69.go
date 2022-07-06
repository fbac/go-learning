/*
	callback: pass a func into a func
*/
package exudemy

import "fmt"

// run a "func(i int) int"
// with the int we pass
func myFuncRunner(x int, f func(i int) int) int {
	return f(x)
}

// return a func(i int) int
func sumN() func(i int) int {
	return func(n int) int {
		funcSum := 0
		for i := 0; i < n; i++ {
			funcSum += i
		}
		return funcSum
	}
}

func main() {
	// call myFuncRunner with an integer and our sumN func (same signature)
	fmt.Println(myFuncRunner(5, sumN()))
}
