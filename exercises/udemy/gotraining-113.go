/*
	Create a customErr struct
*/
package main

import (
	"fmt"
)

type customErr struct {
	id   int
	data string
}

func (err customErr) Error() string {
	return fmt.Sprintf("Error code:\t%v \nReason:\t", err.id) + err.data
}

func NewCustomErr(id int, data string) *customErr {
	return &customErr{id, data}
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	myError := NewCustomErr(1001, "Out of Service")
	foo(myError)
}
