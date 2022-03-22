package repository

import "github.com/kasrasaeed/news_app/entity"

type Repository interface {
	FindAllUsers() ([]*entity.User, error)
}
