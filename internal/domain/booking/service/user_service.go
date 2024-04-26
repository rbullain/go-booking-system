package service

import "go-booking-system/internal/domain/booking/entity"

type UserService interface {
	CreateUser(name string, balance float64) (*entity.UserEntity, error)
}
