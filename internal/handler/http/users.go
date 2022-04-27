package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/altuxa/real-time-forum/internal/model"
)

func (h *Handler) NewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
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
		http.Error(w, "db error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
