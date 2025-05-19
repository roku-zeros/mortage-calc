package server

import "net/http"

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", s.Ping)
	mux.HandleFunc("POST /execute", s.Execute)
	mux.HandleFunc("GET /cache", s.Cache)
}
