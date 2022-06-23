package main

import (
	"fmt"
	"log"
	"os"
)

const (
	aboutMsg = `
	# interfaces io.Writer & io.Reader
	https://pkg.go.dev/io#Writer
	https://pkg.go.dev/io#Reader

	## io.Writer signature
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	
	## io.Reader signature
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	### info
	Many packages that performs some kind of I/O
	implement the interfaces Writer and Reader.
	This makes easy the []byte creation and modification,
	and receive/send it to different I/O

	os.Stdout and os.Stderr are io.Writer type

	### examples
	In package json, the types Encoder and Decoder,
	are methods for io.Writer, such as:
	func NewEncoder(w io.Writer) *Encoder

	Hence, we can open a file (type io.Writer), and 
	create a new encoder/decoder to encode/decode data
	from or to it.

	### low level example
	fmt.Println returns "return Fprintln(os.Stdout, a...)"

	|||
	vvv

	Fprintln signature is "Fprintln(w io.Writer, a ...interface{})"
	
	Hence, os.Stdout is type io.Writer and we can run Fprintln to files
	
	In the package "os" we have the confirmation:
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	`
)

func main() {
	fmt.Println(aboutMsg)

	// write to an io.Writer (os.Stdout)
	// no need to open/close files
	// it already exists and the go implementation
	// manages everything for us
	fmt.Fprintln(os.Stdout, "fmt.Fprintln: writing to a io.Writer (os.Stdout)")

	// create a temp file
	f, err := os.CreateTemp(".", "ioWriter-temp-")
	if err != nil {
		log.Fatalf("os.CreateTemp err: %v\n", err)
	}

	// defer a func to close the file
	// use this anon func to control the returned err
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("f.Close err: %v\n", err)
		}
	}()

	// print to the file using Fprintln
	fmt.Fprintln(f, "fmt.Fprintln: writing to a io.Writer")

	//f2 := os.OpenFile("ioWriter-examples.txt")
}
