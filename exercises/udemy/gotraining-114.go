/*
	Use sqrt.Error struct as value of type error:

	package main

	import (
		"fmt"
		"log"
	)

	type sqrtError struct {
		lat  string
		long string
		err  error
	}

	func (se sqrtError) Error() string {
		return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
	}

	func main() {
		_, err := sqrt(-10.23)
		if err != nil {
			log.Println(err)
		}
	}

	func sqrt(f float64) (float64, error) {
		if f < 0 {
			// write your error code here
		}
		return 42, nil
	}
*/
package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, sqrtError{"50.2289 N", "99.4656 W", errors.New("something happened")}
	}
	return 42, nil
}
