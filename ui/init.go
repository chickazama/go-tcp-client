package ui

import (
	"log"
	"net"

	"github.com/awesome-gocui/gocui"
	"github.com/chickazama/go-tcp-client/client"
)

var (
	c           *client.Client
	name        string
	inputEditor gocui.EditorFunc
)

const (
	maxBufferLength = 128
)

func init() {
	conn, err := net.Dial("tcp4", "127.0.0.1:4444")
	if err != nil {
		log.Fatal(err.Error())
	}
	c = client.New(conn)
	name = "Matt"
	inputEditor = gocui.EditorFunc(setInputEditor)
}
