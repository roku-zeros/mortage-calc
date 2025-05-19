package storage

import (
	"context"
	"mortage-calc/services/calc/internal/models"
)

func (s *Storage) CreateMortage(ctx context.Context, calculation models.Calculation) {
	s.db.Set(calculation)
}

func (s *Storage) GetAllMortages(ctx context.Context) []models.Calculation {
	id := s.db.GetCurrID()
	calculations := make([]models.Calculation, id)
	for i := range id {
		val, ok := s.db.Get(i)
		if !ok {
			continue
		}
		calc := val.(models.Calculation)
		calc.ID = id
		calculations[i] = calc
	}
	return calculations
}
