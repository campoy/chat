package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go match(conn)
	}
}

var partners = make(chan net.Conn)

func match(conn net.Conn) {
	fmt.Fprintln(conn, "Looking for a partner ...")
	select {
	case partners <- conn:
		// the other goroutine won and we can finish
	case p := <-partners:
		chat(conn, p)
	}
}

func chat(a, b net.Conn) {
	fmt.Fprintln(a, "We found a partner")
	fmt.Fprintln(b, "We found a partner")
	go io.Copy(a, b)
	go io.Copy(b, a)
}
