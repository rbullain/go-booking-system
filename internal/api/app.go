package api

import (
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api/controller"
	"go-booking-system/internal/rabbitmq/client"
	"net/http"
)

type Application struct {
	bookingController controller.BookingController
	rabbitmqClient    client.RabbitMQConnection
}

func NewBookingApplication(bookingController controller.BookingController, rabbitmqClient client.RabbitMQConnection) *Application {
	return &Application{
		bookingController: bookingController,
		rabbitmqClient:    rabbitmqClient,
	}
}

func (api *Application) GetUser(ctx *gin.Context) {
	userDTO, err := api.bookingController.GetUserByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, userDTO)
	}
}

func (api *Application) CreateUser(ctx *gin.Context) {
	userDTO, err := api.bookingController.CreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		// TODO: Send UserCreatedEvent message to RabbitQM
		ctx.JSON(http.StatusCreated, userDTO)
	}
}

func (api *Application) GetRoom(ctx *gin.Context) {
	roomDTO, err := api.bookingController.GetRoomByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, roomDTO)
	}
}

func (api *Application) CreateRoom(ctx *gin.Context) {
	roomDTO, err := api.bookingController.CreateRoom(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, roomDTO)
	}
}

func (api *Application) GetReservation(ctx *gin.Context) {
	reservationDTO, err := api.bookingController.GetReservationByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, reservationDTO)
	}
}

func (api *Application) CreateReservation(ctx *gin.Context) {
	reservationDTO, err := api.bookingController.CreateReservation(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		// TODO: Send ReservationCreatedEvent message to RabbitQM
		ctx.JSON(http.StatusCreated, reservationDTO)
	}
}
