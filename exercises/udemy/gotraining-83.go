/*
	Encode json
	send it to os.Stdout
*/
package exudemy

import (
	"encoding/json"
	"log"
	"os"
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

	err := json.NewEncoder(os.Stdout).Encode(potions)
	if err != nil {
		log.Fatalf("err:%v\n", err)
	}
}
