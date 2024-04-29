package rabbitmq

import "encoding/json"

type IPayload interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

type JsonPayload struct {
}

func (payload *JsonPayload) Encode() ([]byte, error) {
	return json.Marshal(payload)
}

func (payload *JsonPayload) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, payload)
}
