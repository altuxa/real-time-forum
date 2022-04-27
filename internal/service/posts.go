package service

import "github.com/altuxa/real-time-forum/internal/repository"

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{
		repo: repo,
	}
}
