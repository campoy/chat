package main

import (
	"fmt"
	"io"
	"log"
)

var partners = make(chan io.ReadWriteCloser)

func match(conn io.ReadWriteCloser) {
	fmt.Fprintln(conn, "Looking for a partner ...")
	select {
	case partners <- conn:
		// the other goroutine won and we can finish
	case p := <-partners:
		chat(conn, p)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "We found a partner")
	fmt.Fprintln(b, "We found a partner")
	errc := make(chan error, 1)
	go copy(a, b, errc)
	go copy(b, a, errc)
	if err := <-errc; err != nil {
		log.Printf("Error chatting: %v", err)
	}
	a.Close()
	b.Close()
}

func copy(a, b io.ReadWriteCloser, errc chan<- error) {
	_, err := io.Copy(a, b)
	errc <- err
}
