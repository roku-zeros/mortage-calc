package providers

import (
	"context"
	"mortage-calc/services/calc/internal/models"
)

type Storage interface {
	CreateMortage(ctx context.Context, calculation models.Calculation)
	GetAllMortages(ctx context.Context) []models.Calculation
}

type MortageProvider struct {
	storage  Storage
}

func NewMortageProvider(storage Storage) MortageProvider {
	return MortageProvider{
		storage:  storage,
	}
}
