package main

import (
	"go-booking-system/cmd/apifx"
	"go-booking-system/cmd/configfx"
	"go-booking-system/cmd/rabbitmqfx"
	"go-booking-system/cmd/serverfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		configfx.Module,
		rabbitmqfx.Module,
		apifx.Module,
		serverfx.Module,
	)
	app.Run()
}
