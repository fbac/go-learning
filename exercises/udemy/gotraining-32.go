// print every char of the uppercase alphabet three times
package exudemy

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U\n", i)
	}
}
