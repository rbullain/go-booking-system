package entity

type IBookingRepository interface {
	IReservationRepository
	IUserRepository
	IRoomRepository
}
