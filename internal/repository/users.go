package repository

import (
	"database/sql"

	"github.com/altuxa/real-time-forum/internal/model"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (r *UsersRepository) Create(user model.User) error {
	stmt, err := r.db.Prepare("INSERT INTO Users (login, email,password) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Login, user.Email, user.Password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
