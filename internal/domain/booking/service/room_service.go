package service

import "go-booking-system/internal/domain/booking/entity"

type RoomService interface {
	CreateRoom(name string, price float64) (*entity.RoomEntity, error)
}
