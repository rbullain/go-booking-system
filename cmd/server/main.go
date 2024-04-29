package main

import (
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api"
	"go-booking-system/internal/api/controller"
	"go-booking-system/internal/application"
	"go-booking-system/internal/infra/repository"
	"go-booking-system/internal/rabbitmq/client"
)

var (
	rabbitClient = client.NewRabbitMQConnection("guest", "guest", "localhost", "5672", "")

	bookingRepository  = repository.NewDatabaseBookingRepository("admin", "admin", "localhost", "3306", "golang_learning_booking")
	bookingService     = application.NewBookingService(bookingRepository)
	bookingController  = controller.NewBookingController(bookingService, bookingService, bookingService)
	bookingApplication = api.NewBookingApplication(bookingController, rabbitClient)
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

	reservationRouter := router.Group("/reservation")
	{
		reservationRouter.POST("/", bookingApplication.CreateReservation)
		reservationRouter.GET("/:id", bookingApplication.GetReservation)
	}

	_ = router.Run("localhost:8080")
}
