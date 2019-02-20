package main

import (
	"io"

	"golang.org/x/net/websocket"
)

type socket struct {
	io.ReadWriteCloser
	done chan bool
}

func (s *socket) Close() error {
	s.done <- true
	return nil
}

func socketHandler(conn *websocket.Conn) {
	s := socket{conn, make(chan bool)}
	go match(&s)
	<-s.done
}
