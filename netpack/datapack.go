package netpack

import (
	"encoding/binary"
)

type DataPack interface {
	Pack(msg Message) []byte
	UnPack(data []byte) Message
}

type DataPackImpl struct {
}

func NewDataPack() DataPack {
	return &DataPackImpl{}
}

func (m *DataPackImpl) Pack(msg Message) []byte {
	data := make([]byte, 4+msg.GetHeadLen())
	binary.BigEndian.PutUint32(data, msg.GetHeadLen())
	copy(data[4:], msg.GetData())
	return data
}

func (m *DataPackImpl) UnPack(data []byte) Message {
	len := binary.BigEndian.Uint32(data)
	return NewClientMsg(data[4 : 4+len])
}
