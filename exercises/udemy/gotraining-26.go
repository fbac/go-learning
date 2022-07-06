package exudemy

import "fmt"

const (
	y2019 = 2019 + iota
	y2020 = 2019 + iota
	y2021 = 2019 + iota
	y2022 = 2019 + iota
)

func main() {
	fmt.Println(y2022)
	fmt.Println(y2021)
	fmt.Println(y2020)
	fmt.Println(y2019)
}
