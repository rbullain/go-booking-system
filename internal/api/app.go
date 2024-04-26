package api

import (
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api/controller"
	"net/http"
)

type Application struct {
	bookingController controller.BookingController
}

func NewBookingApplication(bookingController controller.BookingController) *Application {
	return &Application{
		bookingController: bookingController,
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
