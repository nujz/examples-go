package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(rw http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		rw.WriteHeader(400)
		return
	}
	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("ReadMessage", err)
			break
		}
		if err := c.WriteMessage(mt, msg); err != nil {
			log.Println("WriteMessage", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
