package http

import (
	"log"
	"net/http"

	"github.com/altuxa/real-time-forum/internal/service"
)

type Handler struct {
	usersService        service.Users
	postsService        service.Posts
	commentsService     service.Comments
	categoriesService   service.Categories
	notificationService service.Notification
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		usersService:        services.Users,
		postsService:        services.Posts,
		commentsService:     services.Comments,
		categoriesService:   services.Categories,
		notificationService: services.Notification,
	}
}

func (h *Handler) NewServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.HomePage)
	log.Println("Server started")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalln(err)
	}
}
