package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func httprequest(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("%v", r)
	io.WriteString(w, s)
	r.Body.Close()
}

func httproot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "/")
}

func main() {
	http.HandleFunc("/request", httprequest)
	http.HandleFunc("/", httproot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
