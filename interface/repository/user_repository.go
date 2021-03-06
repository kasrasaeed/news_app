package repository

import "github.com/kasrasaeed/news_app/entity"

type userRepository struct {
	Repository
}

func NewUserRepo(repo Repository) userRepository {
	return userRepository{repo}
}

func (ur *userRepository) FindAllUsers() ([]*entity.User, error) {
	users, err := ur.Repository.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
