package main

import (
	"fmt"
	"net/http"
)

type Mux struct {
	// some import data
}

func (m Mux) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/a" {
		rw.Header().Add("X-Custom-A", "Test-1")
		rw.Write([]byte("123"))
	} else if r.URL.Path == "/b" {
		rw.WriteHeader(401)
		rw.Header().Add("X-Custom-A", "Test-1")
		rw.Write([]byte("123"))
	}
}

func main() {
	mux := Mux{}
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
