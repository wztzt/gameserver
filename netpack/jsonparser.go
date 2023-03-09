package netpack

import (
	"encoding/json"
)

type jsonPaser struct {
}

func (p *jsonPaser) Marshal(msg interface{}) ([]byte, error) {
	return json.Marshal(msg)
}

func (p *jsonPaser) UnMarshal(data []byte, msg interface{}) error {
	return json.Unmarshal(data, msg)
}
