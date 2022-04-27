package app

import (
	"database/sql"
	"log"

	"github.com/altuxa/real-time-forum/internal/handler/http"
	"github.com/altuxa/real-time-forum/internal/repository"
	"github.com/altuxa/real-time-forum/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

func Run() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalln(err)
	}
	repo := repository.NewRepositories(db)
	service := service.NewService(repo)
	handler := http.NewHandler(service)
	handler.NewServer()
}
