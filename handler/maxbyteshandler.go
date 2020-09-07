package handler

import (
	"net/http"

	"github.com/gofaith/rest/internals"
)

func MaxBytesHandler(n int64) func(http.Handler) http.Handler {
	if n <= 0 {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > n {
				internals.Errorf(r, "request entity too large, limit is %d, but got %d, rejected with code %d",
					n, r.ContentLength, http.StatusRequestEntityTooLarge)
				w.WriteHeader(http.StatusRequestEntityTooLarge)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
