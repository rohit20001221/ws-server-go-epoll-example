package server

import (
	"log"

	"golang.org/x/net/websocket"
)

func (s *Server) Broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Println(err)
			}
		}(ws)
	}
}
