package service

import (
	"time"

	"github.com/altuxa/real-time-forum/internal/model"
	"github.com/altuxa/real-time-forum/internal/repository"
	"github.com/altuxa/real-time-forum/jwt"
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

func (s *UsersService) SignIn(user model.User) (string, error) {
	user, err := s.repo.GetByCredentials(user.NickName, user.Password)
	if err != nil {
		return "", err
	}
	jwt := jwt.NewJWT()
	timeExp := time.Now().Add(5 * time.Minute)
	unixExp := timeExp.Unix()
	jwt.NewPayload("user", user.NickName, int(unixExp))
	token, err := jwt.SignedToken("aboba")
	if err != nil {
		return "", err
	}
	return token, nil
}
