package httpserver

import (
	"net/http"
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before
		t1 := time.Now()
		srw := NewStatusResponseWriter(w)
		next.ServeHTTP(srw, r)
		// After
		t2 := time.Now()
		logger.Info(r.RemoteAddr, r.Method, r.RequestURI, r.Proto, srw.statusCode, t2.Sub(t1), r.UserAgent())
	})
}

func NewStatusResponseWriter(w http.ResponseWriter) *StatusResponseWriter {
	return &StatusResponseWriter{
		statusCode:     http.StatusOK,
		ResponseWriter: w,
	}
}

func (srw *StatusResponseWriter) WriteHeader(code int) {
	srw.statusCode = code
	srw.ResponseWriter.WriteHeader(code)
}
