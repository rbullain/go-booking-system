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

	userRoute := router.Group("/user")
	{
		userRoute.POST("/", bookingApplication.CreateUser)
		userRoute.GET("/:id", bookingApplication.GetUser)
	}

	roomRouter := router.Group("/room")
	{
		roomRouter.POST("/", bookingApplication.CreateRoom)
		roomRouter.GET("/:id", bookingApplication.GetRoom)
	}

	_ = router.Run("localhost:8080")
}
