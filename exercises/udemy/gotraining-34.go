package main

import (
	"fmt"
	"time"
)

func main() {

	birthday := 1985

	for {

		fmt.Println(birthday)

		if birthday == time.Now().Year() {
			break
		}

		birthday++
	}

}
