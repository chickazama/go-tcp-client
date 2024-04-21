package main

import (
	"log"
	"net"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:4444")
	if err != nil {
		log.Fatal(err.Error())
	}
	c := New(conn)
	wg.Add(2)
	go c.Send()
	go c.Receive()
	wg.Wait()
}
