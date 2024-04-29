package application

import (
	"go-booking-system/internal/domain/booking/entity"
	"go-booking-system/internal/domain/booking/service"
)

type IBookingService interface {
	service.IUserService
	service.IRoomService
	service.IReservationService
}

type bookingService struct {
	repository entity.IBookingRepository
}

var _ IBookingService = bookingService{}

func NewBookingService(repository entity.IBookingRepository) IBookingService {
	return &bookingService{
		repository: repository,
	}
}

func (service bookingService) GetUserByID(id int64) (*entity.UserEntity, error) {
	user, err := service.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (service bookingService) CreateUser(name string, balance float64) (*entity.UserEntity, error) {
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

func (service bookingService) GetRoomByID(id int64) (*entity.RoomEntity, error) {
	room, err := service.repository.GetRoomByID(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (service bookingService) CreateRoom(name string, price float64) (*entity.RoomEntity, error) {
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

func (service bookingService) CreateReservation(userId int64, roomId int64) (*entity.ReservationEntity, error) {
	newReservation := &entity.ReservationEntity{
		UserID: userId,
		RoomID: roomId,
	}
	newReservation, err := service.repository.CreateReservation(newReservation)
	if err != nil {
		return nil, err
	}
	return newReservation, nil
}

func (service bookingService) GetReservationByID(id int64) (*entity.ReservationEntity, error) {
	reservation, err := service.repository.GetReservationByID(id)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}
