package network

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

//表示链接

type connection interface {
	Start()
	Stop()
}

type connectionImpl struct {
	server   Server   //属于哪个server
	conn     net.Conn //链接
	id       uint32   //id
	msgChan  chan []byte
	exitChan chan bool
}

func NewConnction(server Server, conn net.Conn, id uint32) connection {
	return &connectionImpl{
		server:   server,
		conn:     conn,
		id:       id,
		msgChan:  make(chan []byte),
		exitChan: make(chan bool, 1),
	}
}

func (c *connectionImpl) Start() {
	go c.startReader()
	go c.startWriter()
	c.server.OnConnectionOpen(c)
	log.Println("start ", c.id)

}

func (c *connectionImpl) Stop() {
	c.exitChan <- true

	close(c.exitChan)
	close(c.msgChan)
	log.Println("stop ", c.id)
}

func (c *connectionImpl) startWriter() {
	defer log.Println("writer exit ", c.id)
	for {
		select {
		case data := <-c.msgChan:
			c.conn.Write(data)
		case <-c.exitChan:
			return

		}
	}
}

func (c *connectionImpl) startReader() {
	defer log.Println("reader exit", c.id)
	defer c.Stop()
	for {
		msgLenBuf := make([]byte, 4)
		if _, err := io.ReadFull(c.conn, msgLenBuf); err != nil {
			return
		}
		msgLen := binary.BigEndian.Uint32(msgLenBuf)
		msgData := make([]byte, msgLen)
		if _, err := io.ReadFull(c.conn, msgData); err != nil {
			return
		}
		log.Println(string(msgData), c.id)
		c.SendMsg(msgData)
	}
}

func (c *connectionImpl) SendMsg(data []byte) {

	msgLen := len(data)
	msgData := make([]byte, msgLen+4)
	binary.BigEndian.PutUint32(msgData, uint32(msgLen))
	copy(msgData[4:], data)
	c.msgChan <- msgData
}
