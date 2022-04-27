package service

import (
	"github.com/altuxa/real-time-forum/internal/model"
	"github.com/altuxa/real-time-forum/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) SignUp(user model.User) error {
	err := s.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
