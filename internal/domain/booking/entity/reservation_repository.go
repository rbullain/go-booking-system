package entity

type ReservationRepository interface {
	CreateReservation(reservation *ReservationEntity) (*ReservationEntity, error)
	GetReservationByID(id int64) (*ReservationEntity, error)
}
