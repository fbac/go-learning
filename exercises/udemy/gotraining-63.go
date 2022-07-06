/*
	Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
package exudemy

import "fmt"

func f2() {
	fmt.Println("2nd function - normal execution")
}

func f3() {
	fmt.Println("3rd function - normal execution")
}

func main() {

	defer func() {
		fmt.Println("1st function - deferred")
	}()

	f2()
	f3()

}
