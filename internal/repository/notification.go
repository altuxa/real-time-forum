package repository

import "database/sql"

type NotificationRepository struct {
	db *sql.DB
}

func NewNotifRepo(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}
