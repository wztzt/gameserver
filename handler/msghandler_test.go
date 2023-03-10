package handler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wztzt/gameserver/netpack"

	"github.com/wztzt/gameserver/message"
)

type handle_method func(any interface{})

func handle_ABCD(req message.LoginRequest) {

}

func TestMethod(t *testing.T) {
	method := handle_ABCD
	m_type := reflect.TypeOf(method)
	fmt.Println(m_type)
	fmt.Println(m_type.Kind())

	fmt.Println(m_type.In(1))
	fmt.Println(m_type.In(1).Kind())
	m := reflect.New(m_type.In(1))
	i := m.Interface()
	fmt.Println(m.Kind())
	i.(*message.LoginRequest).Account = "123"
	parser := &netpack.ProtoParser{}
	data, _ := parser.Marshal(i)
	n := reflect.New(m_type.In(1))
	j := n.Interface()
	parser.UnMarshal(data, j)
	//method(1, *j)
	for i := 0; i < m.Elem().NumMethod(); i++ {
		fmt.Println(m.Elem().Method(i))
	}
	fmt.Println(m.Kind())
	m_value := reflect.ValueOf(method)
	fmt.Println(m_value.Kind())
}

func TestMsgHandlerManager(t *testing.T) {
	mgr := DefaultProtoHandlerManager()
	mgr.RegisterHandler(1, handle_LoginRequest)
	req := &message.LoginRequest{
		Account:  "wzert",
		Password: "12334",
	}
	parse := netpack.ProtoParser{}
	data, _ := parse.Marshal(req)
	mgr.HandleMsg(1, data)
}
