package entity

type IRoomRepository interface {
	CreateRoom(room *RoomEntity) (*RoomEntity, error)
	GetRoomByID(id int64) (*RoomEntity, error)
}
