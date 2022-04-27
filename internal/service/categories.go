package service

import "github.com/altuxa/real-time-forum/internal/repository"

type CategoriesService struct {
	repo repository.Categories
}

func NewCategoriesService(repo repository.Categories) *CategoriesService {
	return &CategoriesService{
		repo: repo,
	}
}
