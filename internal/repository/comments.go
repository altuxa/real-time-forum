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

func (c *CommentsRepository) CreateComment(com model.Comment) error {
	stmt, err := c.db.Prepare("INSERT INTO Comments(AuthorID,PostID,BodyText) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(com.AuthodID, com.PostId, com.BodyText)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
