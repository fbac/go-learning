/*
	Create a Encoder that encodes certain struct
	into a file, in format json

	Use io.Writer interface for that purpose
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	fileED string = "encode-decode.txt"
)

/*
	definition of type
	fields have to be public, hence the uppercase vars
*/
type potion struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Effect string `json:"Effect,omitempty"`
	Weight int    `json:"Weight"`
}

func main() {
	// open the new file as RDWR or create it
	fEncoder, err := os.OpenFile(fileED, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("os.Openfile:\t%s\n", err)
	}

	// use the new *os.File to create a Encoder
	// os.File implements io.Writer, so we can
	// create a NewEncoder that will output directly
	// to the file
	encoder := json.NewEncoder((fEncoder))

	// create items
	p1 := potion{
		Id:     0,
		Name:   "lesser health",
		Effect: "you feel slightly better",
		Weight: 1,
	}

	p2 := potion{
		Id:     1,
		Name:   "lesser agility",
		Effect: "you feel a bit more agile",
		Weight: 2,
	}

	p3 := potion{
		Id:     2,
		Name:   "lesser strength",
		Effect: "you feel a bit stronger",
		Weight: 2,
	}

	// create a slice with items
	inventory := []potion{p1, p2, p3}

	// encode the slice inventory
	// it will be written in the encoder output (os.File)
	encoder.Encode(inventory)
}
