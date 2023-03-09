package netpack

import (
	proto "google.golang.org/protobuf/proto"
)

type ProtoParser struct {
}

func (p *ProtoParser) Marshal(msg interface{}) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (p *ProtoParser) UnMarshal(data []byte, msg interface{}) error {
	return proto.Unmarshal(data, msg.(proto.Message))
}
