package dto

type RoomCreateRequestDTO struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type RoomCreateResponseDTO struct {
	ID    int64   `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type RoomRetrieveResponseDTO struct {
	ID    int64   `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
