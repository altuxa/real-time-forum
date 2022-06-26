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
	Tags     []string
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
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
	id := 1 // BETA
	err = h.postsService.CreatePost(model.Post{
		AuthodID: id,
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
