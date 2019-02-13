package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	c := make(chan int, 10)
	go fib(c)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
		time.Sleep(time.Second)
		log.Printf("fib received")
	}
}

func fib(c chan<- int) {
	a, b := 0, 1
	for {
		a, b = b, a+b
		c <- a
		log.Printf("fib sent")
	}
}
