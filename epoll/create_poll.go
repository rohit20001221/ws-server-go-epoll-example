package epoll

import (
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
)

func CreateEpoll() (*Epoll, error) {
	fd, err := syscall.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	return &Epoll{
		FD:          fd,
		Lock:        &sync.RWMutex{},
		Connections: make(map[int]*websocket.Conn),
	}, nil
}
