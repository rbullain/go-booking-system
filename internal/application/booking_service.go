package application

import (
	"go-booking-system/internal/domain/booking/entity"
	"go-booking-system/internal/domain/booking/service"
)

type BookingService interface {
	service.UserService
	service.RoomService
	service.ReservationService
}

type bookingService struct {
	repository entity.BookingRepository
}

func NewBookingService(repository entity.BookingRepository) BookingService {
	return &bookingService{
		repository: repository,
	}
}

func (service *bookingService) CreateUser(name string, balance float64) (*entity.UserEntity, error) {
	newUser := &entity.UserEntity{
		Name:    name,
		Balance: balance,
	}
	newUser, err := service.repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (service *bookingService) CreateRoom(name string, price float64) (*entity.RoomEntity, error) {
	newRoom := &entity.RoomEntity{
		Name:  name,
		Price: price,
	}
	newRoom, err := service.repository.CreateRoom(newRoom)
	if err != nil {
		return nil, err
	}
	return newRoom, nil
}
