package repository

import (
	"database/sql"

	"github.com/altuxa/real-time-forum/internal/model"
)

type Users interface {
	Create(model.User)
}

type Posts interface{}

type Comments interface{}

type Categories interface{}

type Notification interface{}

type Repositories struct {
	Users
	Posts
	Comments
	Categories
	Notification
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:        NewUserRepo(db),
		Posts:        NewPostsRepo(db),
		Comments:     NewCommentsRepo(db),
		Categories:   NewCategoriesRepo(db),
		Notification: NewNotifRepo(db),
	}
}
