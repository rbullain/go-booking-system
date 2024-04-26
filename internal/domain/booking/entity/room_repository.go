package entity

type RoomRepository interface {
	CreateRoom(room *RoomEntity) error
	GetRoomByID(id int64) (*RoomEntity, error)
}
