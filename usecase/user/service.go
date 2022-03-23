package user

import (
	"github.com/kasrasaeed/news_app/domain/entity"
	"github.com/kasrasaeed/news_app/domain/repository/user"
	"github.com/kasrasaeed/news_app/pkg/id"
	"strings"
	"time"
)

type Repository = user.Repository

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateUser(username, password, role string) (id.UUID, error) {
	u, err := entity.NewUser(role, username, password)
	if err != nil {
		return u.Id, err
	}
	return s.repo.Create(u)
}

func (s *Service) GetUser(id id.UUID) (*entity.User, error) {
	return s.repo.GetById(id)
}

func (s *Service) SearchUser(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.GetAll()
}

func (s *Service) DeleteUser(id id.UUID) error {
	u, err := s.GetUser(id)
	if u != nil {
		return entity.NotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

func (s *Service) UpdateUser(e *entity.User) error {
	err := e.Validate()
	if err != nil {
		return entity.InvalidUserEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

func (s *Service) UpdateLastVisit(e *entity.User) error {
	err := entity.UpdateLastVisit(e)
	if err != nil {
		return err
	}
	return s.repo.Update(e)
}
