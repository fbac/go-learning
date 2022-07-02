/*
	given the json string from the last exercise
	unmarshal it again into a struct
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type potion struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var potions []potion

func main() {
	s := `[{"id":0,"name":"health"},{"id":1,"name":"strength"},{"id":2,"name":"agility"}]`
	fmt.Println("string format:\t", s)

	if err := json.Unmarshal([]byte(s), &potions); err != nil {
		log.Fatalf("err: %v\n", err)
	}
	fmt.Println("json format:\t", potions)

}
