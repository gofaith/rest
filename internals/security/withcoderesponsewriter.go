package security

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"reflect"
)

type WithCodeResponseWriter struct {
	Writer http.ResponseWriter
	Code   int
}

func (w *WithCodeResponseWriter) Header() http.Header {
	return w.Writer.Header()
}

func (w *WithCodeResponseWriter) Write(bytes []byte) (int, error) {
	return w.Writer.Write(bytes)
}

func (w *WithCodeResponseWriter) WriteHeader(code int) {
	w.Writer.WriteHeader(code)
	w.Code = code
}

func (w *WithCodeResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.Writer.(http.Hijacker)
	if ok {
		return h.Hijack()
	}
	panic(fmt.Sprintf("WithCodeResponseWriter.Writer:%s is not a http.Hijacker", reflect.TypeOf(w.Writer).String()))
}
