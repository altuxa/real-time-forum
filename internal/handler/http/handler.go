package http

import (
	"log"
	"net/http"

	"github.com/altuxa/real-time-forum/internal/handler/router"
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
	r := new(router.Router)
	r.Route("GET", "/", h.HomePage)
	r.Route("POST", "/users/sign-up", h.SignUp)
	r.Route("POST", "/posts", h.CreatePost)
	r.Route("POST", "/posts/comments", h.CreateComment)
	log.Println("Server started")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}
