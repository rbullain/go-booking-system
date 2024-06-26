package service

import "go-booking-system/internal/domain/booking/entity"

type IReservationService interface {
	CreateReservation(userId int64, roomId int64) (*entity.ReservationEntity, error)
	GetReservationByID(id int64) (*entity.ReservationEntity, error)
}
