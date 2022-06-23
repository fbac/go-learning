/*
	Create a Decoder that decodes certain json
	reading it from file, and output the var
	in go struct format

	Use io.Writer interface for that purpose
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	fileED string = "encode-decode.txt"
)

/*
	definition of type
	fields have to be public, hence the uppercase vars
	the struct has to be created according
	to the json to be received
*/
type potion struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Effect string `json:"Effect,omitempty"`
	Weight int    `json:"Weight"`
}

var inventory []potion

func main() {
	// read the []byte
	// from the file
	dataBytes, err := os.ReadFile(fileED)
	if err != nil {
		fmt.Printf("os.ReadFile err:\t%s\n", err)
	}

	// use the bytes stream to create a io.Reader
	// needed for the Decoder
	byteStream := strings.NewReader(string(dataBytes))

	// create the NewDecoder
	// it will read all the bytes from file
	decoder := json.NewDecoder(byteStream)

	// decode the data into inventory
	// we pass the mem address to the Decoder
	decoder.Decode(&inventory)

	fmt.Printf("%+v\n", inventory)
}
