package user

import (
	"github.com/kasrasaeed/news_app/domain/entity"
	"github.com/kasrasaeed/news_app/pkg/id"
)

type Reader interface {
	GetById(id id.UUID) (*entity.User, error)
	GetByName(name string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
}

type Writer interface {
	Create(e *entity.User) (id.UUID, error)
	Update(e *entity.User) error
	UpdateLastVisit(e *entity.User) error
	Delete(id id.UUID) error
}

type Repository interface {
	Reader
	Writer
}
