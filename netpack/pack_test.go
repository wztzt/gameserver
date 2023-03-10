package netpack

import (
	"fmt"
	"testing"

	"github.com/wzrtzt/GameServer/message"
)

type JsonTest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestPack(t *testing.T) {
	datapack := NewDataPack()
	Msg1 := NewClientMsg([]byte{'h', 'e', 'l', 'l', 'o'})
	data1 := datapack.Pack(Msg1)
	Msg2 := NewClientMsg([]byte{'w', 'o', 'r', 'l', 'd'})
	data2 := datapack.Pack(Msg2)
	data1 = append(data1, data2...)
	sMsg := datapack.UnPack(data1)
	if Msg1.GetHeadLen() == sMsg.GetHeadLen() {
		fmt.Println(string(sMsg.GetData()))
	}
	var parser MsgParser
	parser = &JsonPaser{}
	json_data, _ := parser.Marshal(JsonTest{Id: 123, Name: "你好你好"})
	t.Log(string(json_data))
	var test JsonTest
	parser.UnMarshal(json_data, &test)

	parser = &ProtoParser{}
	loginReq := &message.LoginRequest{
		Account:  "wzwzwz",
		Password: "123456",
	}
	proto_data, _ := parser.Marshal(loginReq)
	loginReq_ := message.LoginRequest{}
	parser.UnMarshal(proto_data, &loginReq_)
}
