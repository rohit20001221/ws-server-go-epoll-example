package epoll

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Epoll struct {
	FD          int
	Connections map[int]*websocket.Conn
	Lock        *sync.RWMutex
}
