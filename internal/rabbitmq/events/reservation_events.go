package events

import (
	"go-booking-system/internal/rabbitmq"
	"time"
)

type ReservationCreated struct {
	rabbitmq.RabbitMQPayload
	ID          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	RoomId      int64     `json:"room_id"`
	CreatedTime time.Time `json:"created_time"`
}
