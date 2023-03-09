package netpack

type Message interface {
	GetHeadLen() uint32
	GetData() []byte
}

type ClientMsgImpl struct {
	MsgLen uint32
	MsgId  int32
	Data   []byte
}

func NewClientMsg(data []byte) Message {
	return &ClientMsgImpl{
		MsgLen: uint32(len(data)),
		Data:   data,
	}
}

func (m *ClientMsgImpl) GetHeadLen() uint32 {
	return m.MsgLen
}

func (m *ClientMsgImpl) GetData() []byte {
	return m.Data
}
