package main

import (
	"fmt"
	"os"
	"time"
)

func efficient(args []string) float64 {
	start := time.Now()
	for i := 0; i < len(args); i++ {
		fmt.Printf("index: %v\targument:%v\n", i, os.Args[i])
	}
	return time.Since(start).Seconds()
}

func inefficient(args []string) float64 {
	start := time.Now()
	var s string
	var sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	return time.Since(start).Seconds()
}

func main() {
	fmt.Printf("func efficient duration:\t%v\n\n", efficient(os.Args))
	fmt.Printf("func inefficient duration:\t%v\n", inefficient(os.Args))
}
