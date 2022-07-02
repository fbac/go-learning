/*
	Print OS and ARCH
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS, ARCH, MAXPROCS:", runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(1))
}
