package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/altuxa/real-time-forum/internal/model"
)

type createCommentInput struct {
	BodyText string
}

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("body read err %s", err), http.StatusInternalServerError)
		return
	}
	input := createCommentInput{}
	err = json.Unmarshal(reqBody, &input)
	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	err = h.commentsService.CreateComment(model.Comment{BodyText: input.BodyText})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
