package main

import (
	"fmt"
	"os"
)

//!+
func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("index: %v\targument:%v\n", i, os.Args[i])
	}
}
