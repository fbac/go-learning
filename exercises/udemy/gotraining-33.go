// print with a for the years you have been alive
package exudemy

import (
	"fmt"
	"time"
)

func main() {
	birthday := 1985
	for birthday <= time.Now().Year() {
		fmt.Println(birthday)
		birthday++
	}
}
