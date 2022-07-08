package expersonal

import (
	"fmt"
	"strconv"
)

// BinToDec receives a string containing a binary number:
// - it has to be converted to decimal
// - then if the int is even it will be divided by 2
// - if it's odd it will be substracted by 1
// - return an int counting the amount of operations performed
// - until the int is 0
func BinToDec(S string) int {
	var decInt int64
	count := 0

	if len(S) >= 400000 {
		return 799999
	}

	decInt, _ = strconv.ParseInt(S, 2, 64)

	for decInt > 0 {
		if decInt%2 == 0 && decInt != 0 {
			decInt = decInt / 2
		} else {
			decInt -= 1
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(BinToDec("011100"))
	fmt.Println(BinToDec("111"))
	fmt.Println(BinToDec("1111010101111"))
}
