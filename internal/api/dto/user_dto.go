package dto

type UserCreateRequestDTO struct {
	Name    string  `json:"name" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}

type UserCreateResponseDTO struct {
	ID      int64   `json:"id" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}
