package main

import (
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api"
	"go-booking-system/internal/api/controller"
	"go-booking-system/internal/application"
	"go-booking-system/internal/infra/repository"
)

var (
	bookingRepository  = repository.NewDatabaseBookingRepository("admin", "admin", "localhost", "3306", "golang_learning_booking")
	bookingService     = application.NewBookingService(bookingRepository)
	bookingController  = controller.NewBookingController(bookingService, bookingService, bookingService)
	bookingApplication = api.NewBookingApplication(bookingController)
)

func main() {
	router := gin.Default()

	router.POST("/user", bookingApplication.CreateUser)
	router.POST("/room", bookingApplication.CreateRoom)

	_ = router.Run("localhost:8080")
}
