package service

import "go-booking-system/internal/domain/booking/entity"

type IUserService interface {
	CreateUser(name string, balance float64) (*entity.UserEntity, error)
	GetUserByID(id int64) (*entity.UserEntity, error)
}
