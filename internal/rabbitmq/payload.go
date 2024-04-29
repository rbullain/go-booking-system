package rabbitmq

type IRabbitMQPayload interface {
	IRabbitMQPayload()
}

type RabbitMQPayload struct {
}

func (RabbitMQPayload) IRabbitMQPayload() {
}
