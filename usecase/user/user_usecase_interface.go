package user

import (
	"github.com/kasrasaeed/news_app/domain/entity"
	"github.com/kasrasaeed/news_app/pkg/id"
)

type UseCase interface {
	GetUser(id id.UUID) (*entity.User, error)
	SearchUser(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(username, password, role string) (id.UUID, error)
	UpdateUser(e *entity.User) error
	UpdateUserLastVisit(e *entity.User) error
	DeleteUser(id id.UUID) error
}
