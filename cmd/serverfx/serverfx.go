package serverfx

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-booking-system/internal/api"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newGinEngine,
	),
	fx.Invoke(registerRoutes),
	fx.Invoke(startWebServer),
)

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
