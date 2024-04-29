package client

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-booking-system/internal/rabbitmq/events"
)

func (connection BookingMessageClient) HandleUserCreatedEvent(msg amqp.Delivery) {
	var event events.UserCreatedEvent
	err := json.Unmarshal(msg.Body, &event)
	if err != nil {
		return
	}

	fmt.Println("User created:", event.Name)
}

func (connection BookingMessageClient) HandleReservationCreatedEvent(msg amqp.Delivery) {
	var event events.ReservationCreated
	err := json.Unmarshal(msg.Body, &event)
	if err != nil {
		return
	}

	fmt.Println("Reservation created:", event.ID)
}
