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

func (r *UsersRepository) GetUserByID(user model.User) (model.User, error) {
	oneUser := model.User{}
	stmt, err := r.db.Prepare("SELECT ID, login, email, password FROM Users WHERE ID = ?")
	if err != nil {
		return model.User{}, err
	}
	row := stmt.QueryRow(user.Id)
	err = row.Scan(&oneUser.Id, &oneUser.Login, &oneUser.Email, &oneUser.Password)
	if err != nil {
		return model.User{}, err
	}
	return oneUser, nil
}

func (r *UsersRepository) GetByCredentials(userName, password string) (model.User, error) {
	var user model.User
	row := r.db.QueryRow("SELECT ID FROM Users WHERE login = ? AND password = ?", userName, password)
	err := row.Scan(&user.Id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
