package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, res, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/echo", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close() // 网络资源
	fmt.Println("res:", res.Status, res.Body)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println("ReadMessage", err)
				return
			}
			if mt == websocket.TextMessage {
				fmt.Println("text", string(msg))
			} else {
				fmt.Println("other", string(msg))
			}
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() // 内部资源

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				fmt.Println("WriteMessage", err)
				return
			}
		case <-interrupt:
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				panic(err)
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
