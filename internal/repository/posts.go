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

func (p *PostsRepository) CreatePost(post model.Post) error {
	stmt, err := p.db.Prepare("INSERT INTO Posts (AuthorID, Title, BodyText) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	row, err := stmt.Exec(post.AuthodID, post.Title, post.BodyText)
	if err != nil {
		return err
	}
	defer stmt.Close()
	id, err := row.LastInsertId()
	if err != nil {
		return err
	}
	for _, i := range post.Tags {
		stmt, err := p.db.Prepare("INSERT INTO Tags (PostID,Value)VALUES(?,?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(int(id), i)
		if err != nil {
			return err
		}
	}
	return nil
}
