package entity

type UserRepository interface {
	CreateUser(user *UserEntity) (*UserEntity, error)
	GetUserByID(id int64) (*UserEntity, error)
}
