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

func (r *UsersRepository) Create(user model.User) {
}
