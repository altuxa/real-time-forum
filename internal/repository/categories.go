package repository

import "database/sql"

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepo(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}
