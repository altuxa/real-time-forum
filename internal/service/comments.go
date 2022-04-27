package service

import "github.com/altuxa/real-time-forum/internal/repository"

type CommentsService struct {
	repo repository.Comments
}

func NewCommentsService(repo repository.Comments) *CommentsService {
	return &CommentsService{
		repo: repo,
	}
}
