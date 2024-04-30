package apifx

import (
	"go-booking-system/cmd/configfx"
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

func newBookingRepository(cfg *configfx.Config) *repository.DatabaseBookingRepository {
	return repository.NewDatabaseBookingRepository(cfg.DatabaseConfig.Username, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.Host, cfg.DatabaseConfig.Port, cfg.DatabaseConfig.Database)
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
