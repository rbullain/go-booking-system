package controller

import "go-booking-system/internal/domain/booking/service"

type BookingController interface {
}

type bookingController struct {
	service service.BookingService
}
