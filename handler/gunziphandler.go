package handler

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/gofaith/rest/httpx"
)

const gzipEncoding = "gzip"

type gzipWriter struct {
	gw *gzip.Writer
	w  http.ResponseWriter
}

func (g *gzipWriter) Header() http.Header {
	return g.w.Header()
}
func (g *gzipWriter) Write(b []byte) (int, error) {
	return g.gw.Write(b)
}
func (g *gzipWriter) WriteHeader(statusCode int) {
	g.WriteHeader(statusCode)
}

func GunzipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get(httpx.ContentEncoding), gzipEncoding) {
			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			r.Body = reader
		}

		if strings.Contains(r.Header.Get(httpx.AcceptEncoding), gzipEncoding) {
			gw := gzip.NewWriter(w)
			defer gw.Flush()
			defer gw.Close()
			w = &gzipWriter{
				gw: gw,
				w:  w,
			}
		}
		next.ServeHTTP(w, r)
	})
}
