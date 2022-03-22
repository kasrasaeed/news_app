package user

import (
	"github.com/kasrasaeed/news_app/entity"
	"strings"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateUser(username, password, role string) (uint32, error) {
	u, err := entity.NewUser(role, username, password)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(u)
}

func (s *Service) GetUser(id uint32) (*entity.User, error) {
	return s.repo.Get(id)
}

func (s *Service) SearchUser(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.GetAll()
}

func (s *Service) DeleteUser(id uint32) error {
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
