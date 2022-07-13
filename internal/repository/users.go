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
	stmt, err := r.db.Prepare("INSERT INTO Users (NickName,Age,Gender,FirstName,LastName,Email,Password) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.NickName, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func (r *UsersRepository) GetUserByID(userID int) (model.User, error) {
	oneUser := model.User{}
	stmt, err := r.db.Prepare("SELECT * FROM Users WHERE ID = ?")
	if err != nil {
		return model.User{}, err
	}
	row := stmt.QueryRow(userID)
	err = row.Scan(&oneUser.Id, &oneUser.NickName, &oneUser.Age, &oneUser.Gender, &oneUser.FirstName, &oneUser.LastName, &oneUser.Email, &oneUser.Password)
	if err != nil {
		return model.User{}, err
	}
	return oneUser, nil
}

func (r *UsersRepository) GetByCredentials(userName, password string) (model.User, error) {
	var user model.User
	row := r.db.QueryRow("SELECT ID FROM Users WHERE NickName = ? AND Password = ?", userName, password)
	err := row.Scan(&user.Id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *UsersRepository) GetPostsByUserID(userId int) ([]model.Post, error) {
	return nil, nil
}
