/*
	Print OS and ARCH
*/
package exudemy

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS, ARCH, MAXPROCS:", runtime.GOOS, runtime.GOARCH, runtime.GOMAXPROCS(1))
}
