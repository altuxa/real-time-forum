package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/altuxa/real-time-forum/internal/model"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body not read", http.StatusInternalServerError)
		return
	}
	user := model.User{}
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		http.Error(w, "unmarshal error", http.StatusInternalServerError)
		return
	}
	err = h.usersService.SignUp(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

type SingInInput struct {
	NickName string
	Password string
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("body read err %s", err), http.StatusBadRequest)
		return
	}
	input := SingInInput{}
	err = json.Unmarshal(reqBody, &input)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid input %s", err), http.StatusBadRequest)
		return
	}
	token, err := h.usersService.SignIn(model.User{
		NickName: input.NickName,
		Password: input.Password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("user not found %s", err), http.StatusBadRequest)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
}

// func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
// 	strId := strings.TrimPrefix(r.URL.Path,"/")
// }
