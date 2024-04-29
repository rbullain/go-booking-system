package main

import (
	"go-booking-system/cmd/rabbitmq/config"
	"go-booking-system/internal/rabbitmq/client"
)

var (
	rabbitClient = client.NewRabbitMQConnection("guest", "guest", "localhost", "5672", "")
)

func main() {
	rabbitClient.InitializeQueues()

	err := rabbitClient.Subscribe(config.QueueUserCreated, rabbitClient.HandleUserCreatedEvent)
	if err != nil {
		panic(err)
	}
	err = rabbitClient.Subscribe(config.QueueReservationCreated, rabbitClient.HandleReservationCreatedEvent)
	if err != nil {
		panic(err)
	}
}
