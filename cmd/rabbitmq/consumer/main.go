package main

import (
	"go-booking-system/cmd/rabbitmq/config"
	"go-booking-system/internal/rabbitmq/client"
)

var (
	rabbitClient = client.NewRabbitMQConnection("guest", "guest", "localhost", "5672", "")
)

func or[T any](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan T)
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {
	rabbitClient.InitializeQueues()

	ch1 := rabbitClient.Subscribe(config.QueueUserCreated, rabbitClient.HandleUserCreatedEvent)
	ch2 := rabbitClient.Subscribe(config.QueueReservationCreated, rabbitClient.HandleReservationCreatedEvent)

	result := or(ch1, ch2)
	<-result
}
