package events

import (
	"go-booking-system/internal/rabbitmq"
	"time"
)

type UserCreatedEvent struct {
	rabbitmq.JsonPayload
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	CreatedTime time.Time `json:"created_time"`
}
