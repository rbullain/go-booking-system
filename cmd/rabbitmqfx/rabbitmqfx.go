package rabbitmqfx

import (
	"go-booking-system/cmd/configfx"
	"go-booking-system/internal/rabbitmq/client"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(newRabbitMQClient),
)

func newRabbitMQClient(cfg *configfx.Config) client.BookingMessageClient {
	return client.NewRabbitMQConnection(cfg.RabbitMQConfig.Username, cfg.RabbitMQConfig.Password, cfg.RabbitMQConfig.Host, cfg.RabbitMQConfig.Port, cfg.RabbitMQConfig.Vhost)
}
