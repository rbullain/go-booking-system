package rabbitmq

import "encoding/json"

type Payload interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

type PayloadBase struct {
}

func (payload *PayloadBase) Encode() ([]byte, error) {
	return json.Marshal(payload)
}

func (payload *PayloadBase) Decode(bytes []byte) error {
	return json.Unmarshal(bytes, payload)
}
