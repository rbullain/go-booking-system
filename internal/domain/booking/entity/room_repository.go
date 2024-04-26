package entity

type RoomRepository interface {
	CreateRoom(room *RoomEntity) (*RoomEntity, error)
	GetRoomByID(id int64) (*RoomEntity, error)
}
