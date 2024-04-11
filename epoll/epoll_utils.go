package epoll

import (
	"net"
	"syscall"

	"github.com/gorilla/websocket"
)

func (epoll *Epoll) Add(ws *websocket.Conn) error {
	fd := GetWebSocketFd(ws)

	err := syscall.EpollCtl(
		epoll.FD,
		syscall.EPOLL_CTL_ADD,
		fd,
		&syscall.EpollEvent{
			Events: syscall.EPOLLIN | syscall.EPOLLHUP,
			Fd:     int32(fd),
		},
	)

	if err != nil {
		return err
	}

	epoll.Lock.Lock()
	defer epoll.Lock.Unlock()

	epoll.Connections[fd] = ws

	return nil
}

func (epoll *Epoll) Remove(ws *websocket.Conn) error {
	fd := GetWebSocketFd(ws)

	err := syscall.EpollCtl(
		epoll.FD,
		syscall.EPOLL_CTL_DEL,
		fd,
		nil,
	)

	if err != nil {
		return err
	}

	epoll.Lock.Lock()
	defer epoll.Lock.Unlock()

	delete(epoll.Connections, fd)

	return nil
}

func (epoll *Epoll) Wait() ([]*websocket.Conn, error) {
	events := make([]syscall.EpollEvent, 100)
	n, err := syscall.EpollWait(epoll.FD, events, 100)

	if err != nil {
		return nil, err
	}

	epoll.Lock.Lock()
	defer epoll.Lock.Unlock()

	connections := make([]*websocket.Conn, 0)

	for i := 0; i < n; i++ {
		conn := epoll.Connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}

	return connections, nil
}

func GetWebSocketFd(ws *websocket.Conn) int {
	conn := ws.NetConn()
	file, err := conn.(*net.TCPConn).File()
	if err != nil {
		panic(err)
	}

	return int(file.Fd())
}
