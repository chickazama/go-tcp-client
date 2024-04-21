package client

import (
	"bufio"
	"net"
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

// The Send method continuously gets a buffer from the Client's
// Outgoing channel, and writes it to its TCP connection.
func (c *Client) Send() error {
	defer c.Connection.Close()
	for buf := range c.Outgoing {
		_, err := c.Connection.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
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
		c.Incoming <- buf
	}
}
