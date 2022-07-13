package http

import (
	"context"
	"net/http"

	_ "github.com/altuxa/real-time-forum/jwt"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "head", "2134")
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
