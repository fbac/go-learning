package main

import (
	"encoding/json"
	"fmt"
)

const (
	aboutMsg = `
	# json marshal
	doc: https://pkg.go.dev/encoding/json#Marshal
	signature: func Marshal(v any) ([]byte, error)

	> given a struct, marshal it as json
	`
)

type potion struct {
	Id     int    `json:"id"`
	Effect string `json:"effect,omitempty"`
}

func main() {
	fmt.Println(aboutMsg)

	var n potion = potion{
		Id:     4,
		Effect: "random effect",
	}

	fmt.Printf("struct potion:\n%+v\n", n)

	bs, _ := json.Marshal(n)
	fmt.Println("\njson representation:")
	fmt.Println(string(bs))
}
