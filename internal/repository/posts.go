package repository

import (
	"database/sql"

	"github.com/altuxa/real-time-forum/internal/model"
)

type PostsRepository struct {
	db *sql.DB
}

func NewPostsRepo(db *sql.DB) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}

func (p *PostsRepository) CreatePost(model.Post) error {
	return nil
}
