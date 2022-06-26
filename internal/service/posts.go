package service

import (
	"github.com/altuxa/real-time-forum/internal/model"
	"github.com/altuxa/real-time-forum/internal/repository"
)

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{
		repo: repo,
	}
}

func (p *PostsService) CreatePost(post model.Post) error {
	err := p.repo.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}
