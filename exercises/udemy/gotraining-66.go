/*
	Build and use an anonnymous func
*/
package exudemy

import "fmt"

func main() {

	defer func() {
		fmt.Println("deferred execution: anonFunc from main")
	}()

	anonFunc := func() {
		fmt.Println("normal execution: anonFunc from var")
	}

	anonFunc()
}
