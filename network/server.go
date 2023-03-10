package network

import (
	"fmt"
	"log"
	"net"

	"github.com/wztzt/gameserver/handler"
)

type Server interface {
	Start()
	Stop()
	OnConnectionOpen(conn connection)
	HandleMsg(id int32, data []byte)
}

type ServerImpl struct {
	address    string
	handlerMgr handler.HandlerManager
}

func NewServer(address string) Server {
	return &ServerImpl{
		address:    address,
		handlerMgr: handler.DefaultProtoHandlerManager(),
	}
}

func (s *ServerImpl) Start() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("error : ", err)
		return
	}
	var id uint32 = 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error : ", err)
			continue

		}
		connection := s.NewConnction(conn, id)
		id++
		go connection.Start()
	}

}

func (s *ServerImpl) NewConnction(conn net.Conn, id uint32) connection {
	return NewConnction(s, conn, id)
}

func (s *ServerImpl) HandleMsg(id int32, data []byte) {
	s.handlerMgr.HandleMsg(id, data)
}

func (s *ServerImpl) OnConnectionOpen(conn connection) {
	fmt.Println("New Connection Open")
}

func (s *ServerImpl) Stop() {

}
