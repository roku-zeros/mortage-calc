package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestLogger(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	loggedHandler := RequestLogger(handler)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	rr := httptest.NewRecorder()

	loggedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestRequestLogger_StatusCode(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	loggedHandler := RequestLogger(handler)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	rr := httptest.NewRecorder()

	loggedHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}
