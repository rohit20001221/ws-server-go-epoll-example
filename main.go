package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rohot20001221/ws-server/epoll"
)

var epoller *epoll.Epoll

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: websocket.IsWebSocketUpgrade,
		}
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		if err := epoller.Add(conn); err != nil {
			conn.Close()
		}
	})

	var err error
	epoller, err = epoll.CreateEpoll()
	if err != nil {
		panic(err)
	}

	go StartEventLoop()

	http.ListenAndServe(":3000", nil)
}

func StartEventLoop() {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			continue
		}

		for _, conn := range connections {
			if conn == nil {
				break
			}

			_, msg, err := conn.ReadMessage()
			if err != nil {
				epoller.Remove(conn)

				conn.Close()
			} else {
				fmt.Println(string(msg))
			}
		}
	}
}
