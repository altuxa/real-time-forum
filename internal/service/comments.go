package service

import (
	"github.com/altuxa/real-time-forum/internal/model"
	"github.com/altuxa/real-time-forum/internal/repository"
)

type CommentsService struct {
	repo repository.Comments
}

func NewCommentsService(repo repository.Comments) *CommentsService {
	return &CommentsService{
		repo: repo,
	}
}

func (c *CommentsService) CreateComment(com model.Comment) error {
	err := c.repo.CreateComment(com)
	if err != nil {
		return err
	}
	return nil
}
