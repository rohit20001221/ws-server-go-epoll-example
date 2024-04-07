package server

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Println("Read Error:", err)
			continue
		}

		msg := buf[:n]

		s.Broadcast(msg)
	}
}
