package apifx

import (
	"go-booking-system/internal/api"
	"go-booking-system/internal/api/controller"
	"go-booking-system/internal/application"
	"go-booking-system/internal/infra/repository"
	"go-booking-system/internal/rabbitmq/client"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newBookingRepository,
		newBookingService,
		newBookingController,
		newBookingApplication,
	),
)

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
