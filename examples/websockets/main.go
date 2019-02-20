package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(conn *websocket.Conn) {
	var msg string
	fmt.Fscanln(conn, &msg)
	fmt.Println("received:", msg)

	fmt.Fprintln(conn, "hello, websocket")
}
