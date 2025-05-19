package server

import (
	"encoding/json"
	"mortage-calc/services/calc/internal/models"
	"net/http"
)

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Calc service OK")) //nolint
}

func (s *Server) Execute(w http.ResponseWriter, r *http.Request) {
	var req models.Params

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = s.mortageProvider.CreateMortage(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) Cache(w http.ResponseWriter, r *http.Request) {
	calculations, err := s.mortageProvider.GetAllMortages(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(calculations)
}
