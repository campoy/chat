package main

import (
	"log"
	"net/http"

	"github.com/golang/net/websocket"
)

/*
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
*/

const listenAddr = "localhost:8080"

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddr)
}
