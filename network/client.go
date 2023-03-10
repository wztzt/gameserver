package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"runtime/debug"
	"sync"

	"github.com/wztzt/gameserver/handler"
	"github.com/wztzt/gameserver/netpack"
	"google.golang.org/protobuf/proto"
)

type client interface {
	Start()
	Stop()
	SendProtoMsg(id int32, msg proto.Message)
	SendMsg(data []byte)
}

type clientImpl struct {
	address    string
	conn       net.Conn
	handlerMgr handler.HandlerManager
	isClosed   bool
	msgChan    chan []byte
	exitChan   chan bool
	mutext     sync.Mutex
}

func NewClient(address string) client {
	return &clientImpl{
		address:    address,
		handlerMgr: handler.DefaultProtoHandlerManager(),
		isClosed:   true,
		msgChan:    make(chan []byte),
		exitChan:   make(chan bool, 1),
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
		data := make([]byte, 1024)
		_, err := c.conn.Read(data)
		if err != nil {
			return
		}
		fmt.Println(string(data))
	}
}

func (c *clientImpl) stratWriter() {
	defer fmt.Println("exit writer")
	for {
		select {
		case data := <-c.msgChan:
			_, err := c.conn.Write(data)
			if err != nil {
				return
			}
		case <-c.exitChan:
			return
		}
	}
}

func (c *clientImpl) SendProtoMsg(id int32, msg proto.Message) {
	parser := &netpack.ProtoParser{}
	data, _ := parser.Marshal(msg)
	data_ := make([]byte, 8+len(data))
	binary.BigEndian.PutUint32(data_, uint32(len(data)))
	binary.BigEndian.PutUint32(data_[4:8], uint32(id))
	copy(data_[8:], data)
	c.SendMsg(data_)
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
	fmt.Println(string(debug.Stack()))
}
