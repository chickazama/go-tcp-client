package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	Connection net.Conn
	Incoming   chan []byte
	Outgoing   chan []byte
}

func New(conn net.Conn) *Client {
	ret := new(Client)
	ret.Connection = conn
	ret.Incoming = make(chan []byte)
	ret.Outgoing = make(chan []byte)
	return ret
}

func (c *Client) Send() error {
	defer c.Connection.Close()
	br := bufio.NewReader(os.Stdin)
	for {
		buf, err := br.ReadBytes('\n')
		if err != nil {
			return err
		}
		br.Reset(os.Stdin)
		buf[len(buf)-1] = 0
		_, err = c.Connection.Write(buf)
		if err != nil {
			return err
		}
	}
}

func (c *Client) Receive() error {
	defer c.Connection.Close()
	br := bufio.NewReader(c.Connection)
	for {
		buf, err := br.ReadBytes(0)
		if err != nil {
			return err
		}
		br.Reset(c.Connection)
		buf[len(buf)-1] = '\n'
		fmt.Printf("%s", buf)
	}
}
