package main

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

	`
)

func main() {

	

}
