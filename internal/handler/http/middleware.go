package http

import (
	"context"
	"fmt"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Aboba", "1234")
		a := r.Header.Get("Dota")
		b := rw.Header().Get("Dota")
		fmt.Println(a, b)
		ctx := context.WithValue(r.Context(), "head", "2134")
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
