package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api"
	"go-booking-system/internal/api/controller"
	"go-booking-system/internal/application"
	"go-booking-system/internal/infra/repository"
	"go-booking-system/internal/rabbitmq/client"
	"go.uber.org/fx"
)

func newRabbitMQClient() client.BookingMessageClient {
	return client.NewRabbitMQConnection("guest", "guest", "localhost", "5672", "")
}

func newBookingRepository() *repository.DatabaseBookingRepository {
	return repository.NewDatabaseBookingRepository("admin", "admin", "localhost", "3306", "golang_learning_booking")
}

func newBookingService(repo *repository.DatabaseBookingRepository) application.IBookingService {
	return application.NewBookingService(repo)
}

func newBookingController(service application.IBookingService) controller.IBookingController {
	return controller.NewBookingController(service, service, service)
}

func newBookingApplication(controller controller.IBookingController, rabbitClient client.BookingMessageClient) *api.Application {
	return api.NewBookingApplication(controller, rabbitClient)
}

func newGinEngine() *gin.Engine {
	return gin.Default()
}

func registerRoutes(router *gin.Engine, application *api.Application) {
	userRoute := router.Group("/user")
	{
		userRoute.POST("/", application.CreateUser)
		userRoute.GET("/:id", application.GetUser)
	}

	roomRouter := router.Group("/room")
	{
		roomRouter.POST("/", application.CreateRoom)
		roomRouter.GET("/:id", application.GetRoom)
	}

	reservationRouter := router.Group("/reservation")
	{
		reservationRouter.POST("/", application.CreateReservation)
		reservationRouter.GET("/:id", application.GetReservation)
	}
}

func startWebServer(lc fx.Lifecycle, routes *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := routes.Run(":8080")
				if err != nil {
					return
				}
			}()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			newRabbitMQClient,
			newBookingRepository,
			newBookingService,
			newBookingController,
			newBookingApplication,
			newGinEngine,
		),
		fx.Invoke(registerRoutes),
		fx.Invoke(startWebServer),
	)
	app.Run()
}
