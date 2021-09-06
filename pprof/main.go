package main

import (
	"net/http"
	_ "net/http/pprof"
)

/*

go tool pprof http://localhost:6060/debug/pprof/profile

wrk -t 4 -c 1000 -d 5s -s /home/z/Documents/rust/minihttp/wrk.lua http://127.0.0.1:6060/a

TODO
*/
func main() {
	http.HandleFunc("/a", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World"))
	})
	http.ListenAndServe("localhost:6060", nil)
}
