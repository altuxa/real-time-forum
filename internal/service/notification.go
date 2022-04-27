package service

import "github.com/altuxa/real-time-forum/internal/repository"

type NotificationService struct {
	repo repository.Notification
}

func NewNotificationService(repo repository.Notification) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}
