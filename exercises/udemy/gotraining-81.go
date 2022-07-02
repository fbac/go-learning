/*
	create a custom type
	Marshal it in json
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

func main() {
	p1 := potion{Id: 0, Name: "health"}
	p2 := potion{Id: 1, Name: "strength"}
	p3 := potion{Id: 2, Name: "agility"}
	potions := []potion{p1, p2, p3}
	fmt.Println("struct format:", potions)

	potionsJson, err := json.Marshal(&potions)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	fmt.Println(string(potionsJson))

}
