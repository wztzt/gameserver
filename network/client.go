package network

import (
	"fmt"
	"net"
)

type client interface {
	Start()
	Stop()
	SendMsg(data []byte)
}

type clientImpl struct {
	address  string
	conn     net.Conn
	msgChan  chan []byte
	exitChan chan bool
}

func NewClient(address string) client {
	return &clientImpl{
		address:  address,
		msgChan:  make(chan []byte),
		exitChan: make(chan bool, 1),
	}
}

func (c *clientImpl) Start() {
	var err error
	c.conn, err = net.Dial("tcp", c.address)
	if err != nil {
		return
	}
	go c.startReader()
	go c.stratWriter()
}

func (c *clientImpl) startReader() {
	defer c.Stop()
	for {
		data := make([]byte, 1024)
		len, err := c.conn.Read(data)
		if err != nil {
			return
		}
		print(string(data[:len]))
	}
}

func (c *clientImpl) stratWriter() {
	defer fmt.Println("exit writer")
	for {
		select {
		case data := <-c.msgChan:
			c.conn.Write(data)
		case <-c.exitChan:
			return
		}
	}
}

func (c *clientImpl) SendMsg(data []byte) {
	c.msgChan <- data
}

func (c *clientImpl) Stop() {
	c.exitChan <- true
	c.conn.Close()
	close(c.msgChan)
	close(c.exitChan)
}
