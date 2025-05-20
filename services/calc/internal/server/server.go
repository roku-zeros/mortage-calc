package server

import (
	"github.com/roku-zeros/mortage-calc/services/calc/internal/providers"
)

type Server struct {
	mortageProvider providers.MortageProvider
}

func New(mortageProvider providers.MortageProvider) *Server {
	return &Server{
		mortageProvider: mortageProvider,
	}
}
