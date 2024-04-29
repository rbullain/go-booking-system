package service

import "go-booking-system/internal/domain/booking/entity"

type IRoomService interface {
	CreateRoom(name string, price float64) (*entity.RoomEntity, error)
	GetRoomByID(id int64) (*entity.RoomEntity, error)
}
