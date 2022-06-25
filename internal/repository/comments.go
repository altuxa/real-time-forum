package repository

import (
	"database/sql"

	"github.com/altuxa/real-time-forum/internal/model"
)

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepo(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{
		db: db,
	}
}

func (c *CommentsRepository) CreateComment(model.Comment) error {
	return nil
}
