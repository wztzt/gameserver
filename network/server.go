package network

import (
	"fmt"
	"log"
	"net"

	"github.com/wzrtzt/GameServer/netpack"
)

type Server interface {
	netpack.MsgParser
	Start()
	Stop()
	OnConnectionOpen(conn connection)
}

type ServerImpl struct {
	address string
	parser  netpack.MsgParser
}

func NewServer(address string) Server {
	return &ServerImpl{
		address: address,
		parser:  &netpack.ProtoParser{},
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

func (s *ServerImpl) OnConnectionOpen(conn connection) {
	fmt.Println("New Connection Open")
}

func (s *ServerImpl) Marshal(msg interface{}) ([]byte, error) {
	return s.parser.Marshal(msg)
}

func (s *ServerImpl) UnMarshal(data []byte, msg interface{}) error {
	return s.parser.UnMarshal(data, msg)
}

func (s *ServerImpl) Stop() {

}
