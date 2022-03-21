package repository

import "github.com/kasrasaeed/clean-architecture-implementation-golang/entity"

type Repository interface {
	FindAllUsers() ([]*entity.User, error)
}