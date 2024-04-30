package rabbitmqfx

import (
	"go-booking-system/internal/rabbitmq/client"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newRabbitMQClient),
)

func newRabbitMQClient() client.BookingMessageClient {
	return client.NewRabbitMQConnection("guest", "guest", "localhost", "5672", "")
}
