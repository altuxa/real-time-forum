package http

import (
	"net/http"

	"github.com/altuxa/real-time-forum/jwt"
	_ "github.com/altuxa/real-time-forum/jwt"
)

// func Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		err := jwt.CheckTokenFromClient(token, "aboba")
// 		if err != nil {
// 			http.Error(rw, err.Error(), http.StatusUnauthorized)
// 			return
// 		}
// 		next.ServeHTTP(rw, r)
// 	})
// }

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		err := jwt.CheckTokenFromClient(token, "aboba")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	})
}
