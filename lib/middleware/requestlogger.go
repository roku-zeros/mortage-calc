package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}


func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()
		next.ServeHTTP(rw, r)
		duration := time.Since(start).Nanoseconds()
		fmt.Printf("%s status_code: %d, duration: %d ns\n", time.Now().Format("2006/01/02 15:04:05"), rw.statusCode, duration)
	})
}
