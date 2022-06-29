/*
	Basic usage of bcrypto:
	https://pkg.go.dev/golang.org/x/crypto/bcrypt

	Theory:
	http://www.usenix.org/event/usenix99/provos/provos.pdf
	https://en.wikipedia.org/wiki/Bcrypt

*/
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// implement a basic password login
func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: <cmd> <pwd-to-lock> <pwd-to-unlock>\n")
	}

	// os.Args[1] would be the user registration endpoint
	// it would save the byteString/hash somewhere (db, fs, ...)
	// we shouldn't store the password anywhere
	// in a production env we'd probably want to protect server's memory as well
	// TODO: discover password provided in os.Args[1] using gdb/dlv
	bs, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.MinCost)
	if err != nil {
		log.Fatalf("err while generating password: %v\n", err)
	}

	// then, the user would unlock using its own password
	// using some other endpoint
	// we just compare the hash and provided password
	// if err == nil -> success
	if err = bcrypt.CompareHashAndPassword(bs, []byte(os.Args[2])); err != nil {
		log.Fatalf("you have no powers here")
	}

	fmt.Printf("login succesful!\n")
	fmt.Printf("password byteString:\t%v\n", bs)
	fmt.Printf("password hashed:\t%s\n", string(bs))
}
