package benchmarkingtest

import "fmt"

func SumInt(x, y int) int {
	return x + y
}

func ConcatenateString(a, b string) string {
	return a + b
}

func main() {
	z := SumInt(5, 5)
	s := ConcatenateString("Hello", " World!")
	fmt.Println(z, s)
}
