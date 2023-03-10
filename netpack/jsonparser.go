package netpack

import (
	"encoding/json"
)

type JsonPaser struct {
}

func (p *JsonPaser) Marshal(msg interface{}) ([]byte, error) {
	return json.Marshal(msg)
}

func (p *JsonPaser) UnMarshal(data []byte, msg interface{}) error {
	return json.Unmarshal(data, msg)
}
