package client

import "go-booking-system/internal/rabbitmq/events"

func (connection BookingMessageClient) SendUserCreatedEvent(event *events.UserCreatedEvent) error {
	err := connection.PublishOnQueue(
		event,
		"user.created",
	)
	if err != nil {
		return err
	}
	return nil
}

func (connection BookingMessageClient) SendReservationCreatedEvent(event *events.ReservationCreated) error {
	err := connection.PublishOnQueue(
		event,
		"reservation.created",
	)
	if err != nil {
		return err
	}
	return nil
}
