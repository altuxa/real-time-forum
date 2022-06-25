package service

import (
	"github.com/altuxa/real-time-forum/internal/model"
	"github.com/altuxa/real-time-forum/internal/repository"
)

type Users interface {
	SignUp(model.User) error
	SignIn(model.User) error
}

type Posts interface {
	CreatePost(model.Post) error
}

type Comments interface {
	CreateComment(model.Comment) error
}

type Categories interface{}

type Notification interface{}

type Services struct {
	Users
	Posts
	Comments
	Categories
	Notification
}

// hz
type ServiceDeps struct {
	Repo *repository.Repositories
}

func NewService(repo *repository.Repositories) *Services {
	return &Services{
		Users:        NewUsersService(repo.Users),
		Posts:        NewPostsService(repo.Posts),
		Comments:     NewCommentsService(repo.Comments),
		Categories:   NewCategoriesService(repo.Categories),
		Notification: NewNotificationService(repo.Notification),
	}
}
