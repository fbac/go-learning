// modulo 4
package exudemy

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
