package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/altuxa/real-time-forum/internal/model"
)

type createPostInput struct {
	Title    string
	BodyText string
	Tags     []model.Tag
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("body read error %s", err), http.StatusInternalServerError)
		return
	}
	inputPost := createPostInput{}
	err = json.Unmarshal(reqBody, &inputPost)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	err = h.postsService.CreatePost(model.Post{
		Title:    inputPost.Title,
		BodyText: inputPost.BodyText,
		Tags:     inputPost.Tags,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
