package handler

import (
	"fmt"
	"reflect"

	"github.com/wzrtzt/GameServer/netpack"
)

type HandleMethod func(msg interface{})

type HandlerManager interface {
	RegisterHandler(id int32, method interface{})
	HandleMsg(id int32, data []byte)
}

type methodtype struct {
	method  reflect.Value
	msgType reflect.Type
}

type HandlerManagerImpl struct {
	parser      netpack.MsgParser
	methodTypes map[int32]*methodtype
}

func DefaultProtoHandlerManager() HandlerManager {
	mgr := &HandlerManagerImpl{
		parser:      &netpack.ProtoParser{},
		methodTypes: make(map[int32]*methodtype),
	}
	mgr.RegisterHandler(1, handle_LoginRequest)
	return mgr
}

func DefaultJsonHandlerManager() HandlerManager {
	mgr := &HandlerManagerImpl{
		parser:      &netpack.JsonPaser{},
		methodTypes: make(map[int32]*methodtype),
	}

	return mgr
}

func (h *HandlerManagerImpl) RegisterHandler(id int32, method interface{}) {
	if _, ok := h.methodTypes[id]; ok {
		fmt.Printf("id = %v has Registered !", id)
		return
	}
	methodType := reflect.TypeOf(method)
	methodValue := reflect.ValueOf(method)
	if methodValue.Kind() == reflect.Func {
		fmt.Println(methodValue.Kind())
	}

	if methodType.NumIn() != 1 {
		fmt.Println("Register Faild")
		return
	}
	h.methodTypes[id] = &methodtype{
		method:  methodValue,
		msgType: methodType.In(0),
	}
}

func (h *HandlerManagerImpl) HandleMsg(id int32, data []byte) {
	if methodType, ok := h.methodTypes[id]; ok {
		typ := methodType.msgType
		if methodType.msgType.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		msg := reflect.New(typ)
		h.parser.UnMarshal(data, msg.Interface())
		methodType.method.Call([]reflect.Value{msg})
		return
	}

}
