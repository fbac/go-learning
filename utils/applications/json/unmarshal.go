package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	aboutMsg = `
	# json unmarshal
	doc: https://pkg.go.dev/encoding/json#Unmarshal
	signature: func Unmarshal(data []byte, v any) error

	> given a json in []byte, initialize a struct 
	with the same fields as the json
	`
)

type potion struct {
	Id     int    `json:"id"`
	Effect string `json:"effect,omitempty"`
}

func main() {
	fmt.Println(aboutMsg)

	// take a json as an input
	// convert it to []byte
	stringjson := `{"id":4,"effect":"random effect"}`
	bytestring := []byte(stringjson)
	fmt.Printf("[1] json string:\t%s\n", stringjson)
	fmt.Printf("[1] byte string:\t%s\n", bytestring)

	// create a new empty struct
	var newPotion potion
	fmt.Printf("[1] new struct potion:\t%+v\n", newPotion)

	// Unmarshal the []byte
	// pass the bytes to the address
	// where the struct is initialized
	fmt.Println("\n> run json.Unmarshal(bytestring, &newPotion)")
	err := json.Unmarshal(bytestring, &newPotion)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	fmt.Printf("[2] initialized struct:\t%+v\n", newPotion)
}
