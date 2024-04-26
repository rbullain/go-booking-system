package dto

type ReservationCreateRequestDTO struct {
	UserID int64 `json:"user_id" binding:"required"`
	RoomID int64 `json:"room_id" binding:"required"`
}

type ReservationCreateResponseDTO struct {
	ID     int64 `json:"id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
	RoomID int64 `json:"room_id" binding:"required"`
}

type ReservationRetrieveResponseDTO struct {
	ID     int64 `json:"id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
	RoomID int64 `json:"room_id" binding:"required"`
}
