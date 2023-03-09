package network

import (
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	server := NewServer(":8989")
	go server.Start()

	time.Sleep(time.Second * 3)

	/*
		client := NewClient("127.0.0.1:8989")
		client.Start()
		buf := time.Now().Unix()
		data := make([]byte, 8+4)
		binary.BigEndian.PutUint32(data, 8)
		binary.BigEndian.PutUint64(data[4:], uint64(buf))
		client.SendMsg(data)
		time.Sleep(time.Second * 2)
		client.Stop()
	*/

	var conns []net.Conn
	for i := 0; i < 100000; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8989")
		if err != nil {
			continue
		}
		conns = append(conns, conn)
	}
	time.Sleep(time.Second * 5)
	/*
		for _, conn := range conns {
			for i := 0; i < 1000; i++ {
				buf := time.Now().Unix()
				data := make([]byte, 8+4)
				binary.BigEndian.PutUint32(data, 8)
				binary.BigEndian.PutUint64(data[4:], uint64(buf))
				conn.Write(data)
				data_ := make([]byte, 8+4)
				io.ReadFull(conn, data_)
				time_ := binary.BigEndian.Uint64(data_[4:])
				if time.Now().Unix()-int64(time_) > 1 {
					fmt.Println("a")
				}
			}
			conn.Close()
		}
	*/
	for _, conn := range conns {
		conn.Close()
	}
	select {}

}
