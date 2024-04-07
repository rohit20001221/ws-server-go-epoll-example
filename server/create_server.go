package server

import "golang.org/x/net/websocket"

func CreateServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}
