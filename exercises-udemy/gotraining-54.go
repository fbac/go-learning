/*
	Create an use an anonymous struct
*/
package main

import "fmt"

func main() {

	anonymous := struct {
		message string
	}{
		message: "yeah, I'm anon",
	}

	fmt.Println(anonymous.message)

}
