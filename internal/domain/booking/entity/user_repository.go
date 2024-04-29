package entity

type IUserRepository interface {
	CreateUser(user *UserEntity) (*UserEntity, error)
	GetUserByID(id int64) (*UserEntity, error)
}
