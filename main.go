package main

import (
	"net/http"

	"github.com/rohot20001221/ws-server/server"
	"golang.org/x/net/websocket"
)

func main() {
	wsServer := server.CreateServer()

	http.Handle("/ws", websocket.Handler(wsServer.HandleWS))
	http.ListenAndServe(":3000", nil)
}
