package netpack

type MsgParser interface {
	Marshal(msg interface{}) ([]byte, error)
	UnMarshal(data []byte, msg interface{}) error
}
