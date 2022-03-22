package user

import "github.com/kasrasaeed/news_app/entity"

type Reader interface {
	Get(id uint32) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	GetAll() ([]*entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (uint32, error)
	Update(e *entity.User) error
	Delete(id uint32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUser(id uint32) (*entity.User, error)
	SearchUser(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(username, password, role string) (uint32, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id uint32) error
}
