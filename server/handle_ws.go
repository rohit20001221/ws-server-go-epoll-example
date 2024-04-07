package server

import (
	"log"

	"golang.org/x/net/websocket"
)

func (s *Server) HandleWS(ws *websocket.Conn) {
	log.Println("Incomming WS Connection", ws.RemoteAddr())

	s.conns[ws] = true
	s.ReadLoop(ws)
}
