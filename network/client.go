package network

import (
	"fmt"
	"net"
	"runtime/debug"
	"sync"
)

type client interface {
	Start()
	Stop()
	SendMsg(data []byte)
}

type clientImpl struct {
	address  string
	conn     net.Conn
	isClosed bool
	msgChan  chan []byte
	exitChan chan bool
	mutext   sync.Mutex
}

func NewClient(address string) client {
	return &clientImpl{
		address:  address,
		isClosed: true,
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
	c.isClosed = false
	go c.startReader()
	go c.stratWriter()
}

func (c *clientImpl) startReader() {
	defer c.Stop()
	for {
		select {
		case <-c.exitChan:
			return
		}
		data := make([]byte, 1024)
		_, err := c.conn.Read(data)
		if err != nil {
			return
		}
		//print(string(data[:len]))
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
	if c.isClosed {
		return
	}
	c.exitChan <- true
	c.conn.Close()
	close(c.msgChan)
	close(c.exitChan)
	c.isClosed = true
	fmt.Println(debug.Stack())
}
