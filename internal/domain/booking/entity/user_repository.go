package entity

type UserRepository interface {
	CreateUser(user *UserEntity) error
	GetUserByID(id int64) (*UserEntity, error)
}
