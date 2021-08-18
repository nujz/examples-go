package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/a", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("X-Custom-A", "Test-1")
		rw.Write([]byte("123"))
	})
	http.HandleFunc("/b", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(401)
		rw.Header().Add("X-Custom-A", "Test-1")
		rw.Write([]byte("123"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
