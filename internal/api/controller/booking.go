package controller

import (
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api/dto"
	"go-booking-system/internal/domain/booking/service"
)

type BookingController interface {
	CreateUser(ctx *gin.Context) (*dto.UserCreateResponseDTO, error)
	CreateRoom(ctx *gin.Context) (*dto.RoomCreateResponseDTO, error)
}

var _ BookingController = bookingController{}

type bookingController struct {
	userService        service.UserService
	roomService        service.RoomService
	reservationService service.ReservationService
}

func NewBookingController(userService service.UserService, roomService service.RoomService, reservationService service.ReservationService) BookingController {
	return &bookingController{
		userService:        userService,
		roomService:        roomService,
		reservationService: reservationService,
	}
}

func (controller bookingController) CreateUser(ctx *gin.Context) (*dto.UserCreateResponseDTO, error) {
	var userDTO dto.UserCreateRequestDTO

	err := ctx.ShouldBind(&userDTO)
	if err != nil {
		return nil, err
	}

	user, err := controller.userService.CreateUser(userDTO.Name, userDTO.Balance)
	if err != nil {
		return nil, err
	}

	response := &dto.UserCreateResponseDTO{
		ID:      user.ID,
		Name:    user.Name,
		Balance: user.Balance,
	}
	return response, nil
}

func (controller bookingController) CreateRoom(ctx *gin.Context) (*dto.RoomCreateResponseDTO, error) {
	var roomDTO dto.RoomCreateRequestDTO

	err := ctx.ShouldBind(&roomDTO)
	if err != nil {
		return nil, err
	}

	room, err := controller.roomService.CreateRoom(roomDTO.Name, roomDTO.Price)
	if err != nil {
		return nil, err
	}

	response := &dto.RoomCreateResponseDTO{
		ID:    room.ID,
		Name:  room.Name,
		Price: room.Price,
	}
	return response, nil
}
