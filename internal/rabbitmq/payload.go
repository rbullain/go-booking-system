package rabbitmq

type Payload interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}
