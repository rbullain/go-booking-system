package entity

type BookingRepository interface {
	ReservationRepository
	UserRepository
	RoomRepository
}
